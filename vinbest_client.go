package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type vinbestClientOptions struct {
	ServerHost       string
	ServerPort       int
	AESEncryptionKey []byte
	Username         string
	PPKNum           int
	Pwd              string
	LicenseKey       []int
	MessageDecoder   func(message []byte) ([]byte, error)
	KeyPreparer      func(key []byte) ([]byte, error)
	MessageEncoder   func(message []byte) ([]byte, error)
	Logger           *logrus.Logger
}

type vinbestClient struct {
	vinbestClientOptions
}

func NewVinbestClient(options vinbestClientOptions) *vinbestClient {
	return &vinbestClient{
		options,
	}
}

type JSON map[string]interface{}

func (client *vinbestClient) recognizeMessage(message []byte) (interface{}, error) {
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

func (client *vinbestClient) Run() error {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%v", client.ServerHost, client.ServerPort), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	wsLogger := client.Logger.WithField("service", "WS")

	var lastSentMessage []byte

	send := func(data []byte) error {
		lastSentMessage = data
		wsLogger.WithField("data", data).Printf("send message")
		return conn.WriteMessage(websocket.TextMessage, data)
	}

	messages := make(chan []byte)

	encodedKey, err := client.KeyPreparer(client.AESEncryptionKey)
	if err != nil {
		client.Logger.WithError(err).Error("can't prepare key")
		return err
	}

	if err := send(encodedKey); err != nil {
		client.Logger.WithError(err).Error("can't send key")
		return err
	}

	go func() {
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
	}()

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
						logger.WithField("last_sent_message", lastSentMessage).WithField("response_error", val["message"]).Info("got error from WS server")
						break
					case "key_accepted":
						encodedUserData, err := client.prepareUserData()
						if err != nil {
							logger.WithError(err).Error("can't prepare user data")
							return
						}

						if err := send(encodedUserData); err != nil {
							logger.WithError(err).Error("can't send WS message")
							return
						}
						break
					default:
						// {"events":[{"data":[72,16],"time":1577253372112},{"data":[8,40],"time":1577253372112}],"ppk_num":286,"type":"events"}
						logger.Error("can't handle json response")
					}
					break
				case string:
					logger.Debug("Message is string: %+v", val)
				default:
					logger.Debug("default")
				}
			}(message)
		default:
			continue
		}
	}
}

// AES data
func (client *vinbestClient) prepareUserData() ([]byte, error) {
	data := []byte(`{"type":"login","user_name":"a.shtovba","ping_interval":45000,"ppks":[{"ppk_num":286,"pwd":"123456","license_key":[73,10,7,39,4,50]}]}`)

	encoded, err := client.MessageEncoder(data)
	if err != nil {
		return nil, err
	}

	encoded = []byte(`U2FsdGVkX19U+bdUkih6g4ejjpnYYGH2cw0c5LFJ98dczaZv9+KfEzCLPviqQWRFxQR59UP8rHI7zOrZiqs6+ijAHyF0Z34SHquc2dTTXyCwQ9jvyX2tmKJSHWVV+KsR23goCYOyJXA1/kPr9KrV0o5AQLbTza55465BUdMENOP7HiDiVWQRjozDxCdiLHx0MP+Jx1DN21WTkyV7RDVHtw==`)

	client.Logger.WithField("plain", data).WithField("encoded", encoded).Printf("prepare AES encoded user data")

	return encoded, nil
}
