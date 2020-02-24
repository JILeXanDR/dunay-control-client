package venbest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JILeXanDR/dunay-control-client/encryption"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type ClientOptions struct {
	HardCodedLoginData []byte
	RSAPublicKey       string
	ServerHost         string
	ServerPort         int
	Username           string
	PPKNum             uint
	Pwd                string
	LicenseKey         []uint
	Logger             *logrus.Logger
}

type ClientErr struct {
	message string
}

func (err ClientErr) Error() string {
	return err.message
}

func generateKey(length int) []byte {
	// TODO: generate random string when EAS encyption will by fixed
	return []byte("1111111111111111")
}

type Client struct {
	options ClientOptions
	Events  chan PPKEvent
	States  chan State
	Errors  chan ClientErr
	key     []byte
	ws      *ws
	aes     *encryption.AESEncryptionService
	rsa     *encryption.RSAEncryptionService
	logger  *logrus.Logger
}

func NewClient(options ClientOptions) *Client {
	key := generateKey(16)
	//options.Logger.WithField("key", string(key)).Printf("generated AES key")

	return &Client{
		options,
		make(chan PPKEvent),
		make(chan State),
		make(chan ClientErr),
		key,
		newWS(
			&url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%v", options.ServerHost, options.ServerPort)},
			options.Logger.WithField("module", "WS").Logger,
		),
		encryption.NewAESEncryptionService(key),
		encryption.NewRSAEncryptionService([]byte(options.RSAPublicKey)),
		options.Logger,
	}
}

func (client *Client) parseServerResponse(message []byte) (interface{}, error) {
	// for errors
	if json.Valid(message) {
		var res JSON

		err := json.Unmarshal(message, &res)

		return res, err
	}

	plainJSON, err := client.aesDecodeMessage(message)
	if err != nil {
		return nil, err
	}

	client.logger.WithField("value", plainJSON).Printf("Plain JSON")

	var res JSON
	if err := json.Unmarshal(plainJSON, &res); err != nil {
		return nil, err
	}

	if string(message) == "ping" || string(message) == "pong" {
		return string(message), nil
	}

	return nil, errors.New("unknown response")
}

func (client *Client) Start() error {
	closeFunc, err := client.ws.connect()
	if err != nil {
		client.logger.WithError(err).Error("can't connect to WS server")
		return err
	}
	defer closeFunc()

	encodedKey, err := client.rsa.Encode(client.key)
	if err != nil {
		client.logger.WithError(err).Error("can't prepare key")
		return err
	}

	client.logger.Debugf("send encoded public key")
	if err := client.ws.send(encodedKey); err != nil {
		client.logger.WithError(err).Error("can't send key")
		return err
	}

	go client.ws.readMessages()

	for {
		select {
		case message := <-client.ws.processMessage:
			client.logger.Debugf("got message to process")
			client.processSingleMessage(message)
		}
	}
}

// AES data
func (client *Client) loginData() ([]byte, error) {
	// TODO: пока код дальше не используется, данные юзера захардкожены
	// FIXME: написать корректный AES encrypt
	return client.options.HardCodedLoginData, nil
}

func (client *Client) encodeMessage(message []byte) ([]byte, error) {
	b, err := client.aes.Encode(message)
	if err != nil {
		return nil, err
	}

	return []byte(base64.StdEncoding.EncodeToString(b)), nil
}

func (client *Client) aesDecodeMessage(message []byte) ([]byte, error) {
	return client.aes.Decode(message)
}

func (client *Client) processSingleMessage(message []byte) {
	logger := client.logger.WithField("original_message", string(message))
	logger.Debug("start processing message...")

	defer func() {
		logger.Debug("processing message done.")
	}()

	res, err := client.parseServerResponse(message)
	if err != nil {
		logger.WithError(err).Error("can't recognize message")
		return
	}

	logger.WithField("parsed", res).Debug("parsed server response")

	switch val := res.(type) {
	case JSON:
		logger := client.logger.WithField("data", val)

		switch val["type"].(string) {
		case "error":
			client.Errors <- ClientErr{message: val["message"].(string)}
			break
		case "key_accepted":
			encodedLoginData, err := client.loginData()
			if err != nil {
				logger.WithError(err).Error("can't prepare user data")
				return
			}

			if err := client.ws.send(encodedLoginData); err != nil {
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
