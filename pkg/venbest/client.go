package venbest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ClientOptions struct {
	ServerHost     string
	ServerPort     int
	Username       string
	PPKNum         uint
	Pwd            string
	LicenseKey     []uint
	Logger         *logrus.Logger
	MessageDecoder func(message []byte) ([]byte, error)
	KeyPreparer    func() ([]byte, error)
	MessageEncoder func(message []byte) ([]byte, error)
	StateHandler   func(payload map[string]interface{})
	EventsHandler  func(payload map[string]interface{})
}

type Client struct {
	ClientOptions
}

func NewClient(options ClientOptions) *Client {
	return &Client{
		options,
	}
}

type JSON map[string]interface{}

func (client *Client) recognizeMessage(message []byte) (interface{}, error) {
	if string(message) == "ping" || string(message) == "pong" {
		return string(message), nil
	}

	var res JSON

	if err := json.Unmarshal(message, &res); err == nil {
		client.Logger.WithField("value", message).Printf("Plain JSON")
	} else {
		plainJSON, err := client.MessageDecoder(message)
		if err != nil {
			return nil, err
		}

		client.Logger.WithField("value", plainJSON).Printf("Plain JSON")

		if err := json.Unmarshal(plainJSON, &res); err != nil {
			return nil, err
		}
	}

	client.Logger.WithField("value", res).Printf("JSON structure")

	return res, nil
}

func (client *Client) Run() error {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%v", client.ServerHost, client.ServerPort), nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			client.Logger.WithError(err).Error("failed to close WS connection")
		}
	}()

	wsLogger := client.Logger.WithField("service", "WS")

	messages := make(chan []byte)

	encodedKey, err := client.KeyPreparer()
	if err != nil {
		client.Logger.WithError(err).Error("can't prepare key")
		return err
	}

	if err := client.wsSend(wsLogger.Logger, conn, encodedKey); err != nil {
		client.Logger.WithError(err).Error("can't send key")
		return err
	}

	go client.readMessages(conn, wsLogger.Logger, messages)

	client.processMessages(conn, wsLogger.Logger, messages)

	return nil
}

// AES data
func (client *Client) prepareUserData() ([]byte, error) {
	data := loginData{
		Type:         "login",
		UserName:     client.Username,
		PingInterval: 60000,
		PPKs: []PPK{
			{
				PPKNum:     client.PPKNum,
				Pwd:        client.Pwd,
				LicenceKey: client.LicenseKey,
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	println("jsonData", string(jsonData))

	encoded, err := client.MessageEncoder(jsonData)
	if err != nil {
		return nil, err
	}

	// TODO: fix encoder above
	encoded = []byte(`U2FsdGVkX19jIyfUv0bkdP9Kcr1yP83XEyGrEZ0Pdn9aGKWYufmKcOozJtPnB7UyklqanvTIwMkm/V2+foQA3BGlbkAnEhZyEsbyzCW8RjuOZ9M/Ir1aOe+cO+gSIImXPjH8hE/qIlkXB0piOW53AJTO52BS4s7kMOat+RChiDDVlBgZL3ABzAHHb/SU+LC211AARjdooKxnIEA/n/B6PA==`)

	client.Logger.WithField("plain", data).WithField("encoded", string(encoded)).Debug("prepare AES encoded user data")

	return encoded, nil
}

func (client *Client) readMessages(conn *websocket.Conn, wsLogger *logrus.Logger, messages chan []byte) {
	for {
		wsLogger.Println("waiting for incoming message...")
		_, message, err := conn.ReadMessage()
		if err != nil {
			wsLogger.WithError(err).Printf("can't read message")
			continue
		}

		wsLogger.WithField("data", message).Printf("Received plain message")

		messages <- message
	}
}

func (client *Client) wsSend(logger *logrus.Logger, conn *websocket.Conn, data []byte) error {
	logger.WithField("data", data).Printf("send message")
	return conn.WriteMessage(websocket.TextMessage, data)
}

func (client *Client) processMessages(conn *websocket.Conn, logger *logrus.Logger, messages chan []byte) {
	for {
		select {
		case message := <-messages:
			go func(message []byte) {

				logger := client.Logger.WithField("original_message", string(message))

				logger.Debug("start processing message...")

				defer func() {
					logger.Debug("processing message done.")
				}()

				res, err := client.recognizeMessage(message)
				if err != nil {
					logger.WithError(err).Error("can't recognize message")
					return
				}

				switch val := res.(type) {
				case JSON:

					logger := client.Logger.WithField("data", val)
					logger.Debug("recognized message")

					switch val["type"].(string) {
					case "error":
						logger.WithField("response_error", val["message"]).Info("got error from WS server")
						break
					case "key_accepted":
						encodedUserData, err := client.prepareUserData()
						if err != nil {
							logger.WithError(err).Error("can't prepare user data")
							return
						}

						if err := client.wsSend(logger.Logger, conn, encodedUserData); err != nil {
							logger.WithError(err).Error("can't send WS message")
							return
						}
						break
					case "events":
						fields := logrus.Fields{
							"ppk":    val["ppk_num"],
							"events": val["events"],
						}
						logger.WithFields(fields).Debug("got new events")
						client.EventsHandler(val)
						break
					case "state":
						client.StateHandler(val)
						break
					default:
						logger.Error("can't handle json response")
					}
					break
				case string:
					logger.Debugf("Message is string: %+v", val)
					break
				default:
					logger.Debug("default")
				}
			}(message)
		default:
			continue
		}
	}
}
