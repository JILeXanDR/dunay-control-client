package venbest

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

type ClientOptions struct {
	ServerHost      string
	ServerPort      int
	Username        string
	PPKNum          uint
	Pwd             string
	LicenseKey      []uint
	Logger          *logrus.Logger
	EncodeMessage   func(message []byte) ([]byte, error)
	DecodeMessage   func(message []byte) ([]byte, error)
	EncodeKey       func() ([]byte, error)
	PrepareUserData func() ([]byte, error)
}

type Client struct {
	ClientOptions
	Events chan PPKEvent
	States chan State
	ws     *ws
}

func NewClient(options ClientOptions) *Client {
	return &Client{
		options,
		make(chan PPKEvent),
		make(chan State),
		&ws{
			&url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%v", options.ServerHost, options.ServerPort)},
			make(chan []byte),
			nil,
			options.Logger.WithField("service", "WS").Logger,
		},
	}
}

func (client *Client) decodeMessage(message []byte) (interface{}, error) {
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

func (client *Client) Start() error {
	closeFunc, err := client.ws.connect()
	if err != nil {
		client.Logger.WithError(err).Error("can't connect to WS server")
		return err
	}
	defer closeFunc()

	encodedKey, err := client.EncodeKey()
	if err != nil {
		client.Logger.WithError(err).Error("can't prepare key")
		return err
	}

	if err := client.ws.send(encodedKey); err != nil {
		client.Logger.WithError(err).Error("can't send key")
		return err
	}

	go client.ws.readMessages()

	for {
		select {
		case message := <-client.ws.processMessage:
			client.processSingleMessage(message)
		}
	}
}

// AES data
func (client *Client) prepareUserData() ([]byte, error) {
	if client.PrepareUserData != nil {
		return client.PrepareUserData()
	}

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

	client.Logger.WithField("plain", data).WithField("encoded", string(encoded)).Debug("prepare AES encoded user data")

	return encoded, nil
}

func (client *Client) processSingleMessage(message []byte) {
	logger := client.Logger.WithField("original_message", string(message))

	logger.Debug("start processing message...")

	defer func() {
		logger.Debug("processing message done.")
	}()

	res, err := client.decodeMessage(message)
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

			if err := client.ws.send(encodedUserData); err != nil {
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
