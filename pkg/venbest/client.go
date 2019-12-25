package venbest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

type ClientOptions struct {
	ServerHost    string
	ServerPort    int
	Username      string
	PPKNum        uint
	Pwd           string
	LicenseKey    []uint
	Logger        *logrus.Logger
	EncodeMessage func(message []byte) ([]byte, error)
	DecodeMessage func(message []byte) ([]byte, error)
	EncodeKey     func() ([]byte, error)
}

type Client struct {
	ClientOptions
	Events         chan PPKEvent
	States         chan State
	processMessage chan []byte
	// conn represents WS connection with server
	conn *websocket.Conn
}

func NewClient(options ClientOptions) *Client {
	return &Client{
		options,
		make(chan PPKEvent),
		make(chan State),
		make(chan []byte),
		nil,
	}
}

func (client *Client) recognizeMessage(message []byte) (interface{}, error) {
	if string(message) == "ping" || string(message) == "pong" {
		return string(message), nil
	}

	var res JSON

	if err := json.Unmarshal(message, &res); err == nil {
		client.Logger.WithField("value", message).Printf("Plain JSON")
	} else {
		plainJSON, err := client.DecodeMessage(message)
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

	client.conn = conn

	wsLogger := client.Logger.WithField("service", "WS")

	encodedKey, err := client.EncodeKey()
	if err != nil {
		client.Logger.WithError(err).Error("can't prepare key")
		return err
	}

	if err := client.wsSend(wsLogger.Logger, encodedKey); err != nil {
		client.Logger.WithError(err).Error("can't send key")
		return err
	}

	//go func() {
	//	time.Sleep(3 * time.Second)
	//	client.processMessage <- []byte(`{
	//  "events": [
	//    {
	//      "data": [
	//        72,
	//        16
	//      ],
	//      "time": 1577253372112
	//    },
	//    {
	//      "data": [
	//        64,
	//        40
	//      ],
	//      "time": 1577253372112
	//    }
	//  ],
	//  "ppk_num": 286,
	//  "type": "events"
	//}`)
	//}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	client.processMessage <- []byte(`{
	//  "fallbackIP": "193.108.249.060",
	//  "ppks": [
	//    {
	//      "accum": 1,
	//      "door": 1,
	//      "group": 0,
	//      "groups": {
	//        "1": 0
	//      },
	//      "lines": {
	//        "1": 88,
	//        "2": 88,
	//        "3": 88,
	//        "4": 88,
	//        "5": 88,
	//        "6": 88,
	//        "7": 88,
	//        "8": 88
	//      },
	//      "online": 1,
	//      "power": 1,
	//      "ppk_num": 286,
	//      "relay2": 0,
	//      "scenario": [],
	//      "uk2": 0,
	//      "uk3": 0
	//    },
	//    {
	//      "accum": 1,
	//      "door": 1,
	//      "group": 0,
	//      "groups": {
	//        "1": 0
	//      },
	//      "lines": {
	//        "1": 88,
	//        "2": 88,
	//        "3": 88,
	//        "4": 88,
	//        "5": 88,
	//        "6": 88,
	//        "7": 88,
	//        "8": 88
	//      },
	//      "online": 1,
	//      "power": 1,
	//      "ppk_num": 286,
	//      "relay2": 0,
	//      "scenario": [],
	//      "uk2": 0,
	//      "uk3": 0
	//    }
	//  ],
	//  "time": 1577290359942,
	//  "type": "state"
	//}`)
	//}()

	go client.readMessages(wsLogger.Logger)

	for {
		select {
		case message := <-client.processMessage:
			client.processSingleMessage(message)
		}
	}
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

	encoded, err := client.EncodeMessage(jsonData)
	if err != nil {
		return nil, err
	}

	// TODO: fix encoder above
	// a.shtovba
	encoded = []byte(`U2FsdGVkX19jIyfUv0bkdP9Kcr1yP83XEyGrEZ0Pdn9aGKWYufmKcOozJtPnB7UyklqanvTIwMkm/V2+foQA3BGlbkAnEhZyEsbyzCW8RjuOZ9M/Ir1aOe+cO+gSIImXPjH8hE/qIlkXB0piOW53AJTO52BS4s7kMOat+RChiDDVlBgZL3ABzAHHb/SU+LC211AARjdooKxnIEA/n/B6PA==`)

	// a.shovba1
	encoded = []byte(`U2FsdGVkX18uEU7JYCAOV8M4ZRp+OzdSiRsl7XBraCD4u6VBsbsOVeir4XqHCpRKPdcenofKC4vTv5eTVzqoF0E7idf4y30SV3bHTh9okGnkhGVcrE/bJjFSKcU/sqEOXr9T3l3HpFgfEbLzenp5JEEEBEH01v2CnNYDXtqpOIFhgooKddYaParafe0SeV7pAKS1YpMB7lUHJKLtP3JA2w==`)

	client.Logger.WithField("plain", data).WithField("encoded", string(encoded)).Debug("prepare AES encoded user data")

	return encoded, nil
}

func (client *Client) readMessages(wsLogger *logrus.Logger) {
	for {
		wsLogger.Println("waiting for incoming message...")
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			wsLogger.WithError(err).Printf("can't read message")
			continue
		}

		wsLogger.WithField("data", message).Printf("Received plain message")

		client.processMessage <- message
	}
}

func (client *Client) wsSend(logger *logrus.Logger, data []byte) error {
	logger.WithField("data", data).Printf("send message")
	return client.conn.WriteMessage(websocket.TextMessage, data)
}

func (client *Client) processSingleMessage(message []byte) {
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

			if err := client.wsSend(logger.Logger, encodedUserData); err != nil {
				logger.WithError(err).Error("can't send WS message")
				return
			}
			break
		case "events":
			events, ok := val["events"].([]interface{})
			if ok {
				for _, eventVal := range events {
					event, ok := eventVal.(map[string]interface{})
					if ok {
						data, _ := event["data"].([]interface{})
						timestamp := event["time"].(float64)

						if len(data) == 2 {
							client.Events <- PPKEvent{
								Code:       EventCode(data[0].(float64)),
								Additional: uint(data[1].(float64)),
								When:       time.Unix(int64(timestamp/1000), 0),
							}
						}
					}
				}
			}
			break
		case "state":
			state := State{
				PPKs: []PPKState{},
				When: time.Unix(int64(val["time"].(float64)/1000), 0),
			}

			for _, ppk := range val["ppks"].([]interface{}) {
				state.PPKs = append(state.PPKs, PPKState{ppk.(map[string]interface{})})
			}

			client.States <- state
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
}
