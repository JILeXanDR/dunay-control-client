package main

import (
	"dunay-control-client/pkg/encryption"
	"dunay-control-client/pkg/venbest"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

var config Config

func main() {
	if err := readConfig("config.json", &config); err != nil {
		log.Fatal(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: false})
	logger.SetLevel(logrus.TraceLevel)

	// TODO: generate with `genKey(16)` when EAS encyption will by fixed
	key := []byte(config.AESPassword)
	//key = genKey(16)

	logger.WithField("key", string(key)).Printf("generated AES key")

	rsaEncryptionService := encryption.NewRSAEncryptionService([]byte(config.RSAPublicKey))
	aesEncryptionService := encryption.NewAESEncryptionService(key)

	client := venbest.NewClient(venbest.ClientOptions{
		ServerHost: config.VenbestServer,
		ServerPort: config.VenbestPort,
		Username:   config.VenbestUsername,
		PPKNum:     config.VenbestPPKNum,
		Pwd:        config.VenbestPwd,
		LicenseKey: config.VenbestLicenseKey,
		Logger:     logger,
		PrepareUserData: func() ([]byte, error) {
			//{
			//	"type": "login",
			//	"user_name": "andersen_bot",
			//	"ping_interval": 60000,
			//	"ppks": [{ "ppk_num": 286, "pwd": "123456", "license_key": [73, 10, 7, 39, 4, 50] }]
			//  }
			return []byte(`U2FsdGVkX1+pkUSfOyfkR9CRTllynWSQCZln8K/Imm7X76PYnpfaZmBmHyRFXXBFh9BZZFsXEkw2PnOnQJ1PMSPcTpVBQil9mtatUvBZenzSYgO2aRn2ygvcc43pVDLOcOuss/gY7OIUAPix+yotyjmoq0NG8RNuaSvyz3EY7FW6csTuD94WV6Tp9251MFlz+Yg1RYAc0CXcICR4TMQbbQ==`), nil
		},
		EncodeMessage: func(message []byte) ([]byte, error) {
			b, err := aesEncryptionService.Encode(message)
			if err != nil {
				return nil, err
			}

			return []byte(base64.StdEncoding.EncodeToString(b)), nil
		},
		DecodeMessage: func(message []byte) ([]byte, error) {
			return aesEncryptionService.Decode(message)
		},
		EncodeKey: func() ([]byte, error) {
			return rsaEncryptionService.Encode(key)
		},
	})

	go func() {
		for {
			select {
			case event := <-client.Events:
				logger.WithField("info", event).Debug("new event happened")

				switch event.Code {
				case venbest.EventCode64:
					sendSkypeMessage(fmt.Sprintf("Офис закрыт (%s)", event.When.Format(time.RFC1123)))
				case venbest.EventCode72:
					sendSkypeMessage(fmt.Sprintf("Офис открыт (%s)", event.When.Format(time.RFC1123)))
				case venbest.EventCode108:
					sendSkypeMessage(fmt.Sprintf("Открыта дверца ППК (%s)", event.When.Format(time.RFC1123)))
				case venbest.EventCode109:
					sendSkypeMessage(fmt.Sprintf("Закрыта дверца ППК (%s)", event.When.Format(time.RFC1123)))
				default:
					//send(fmt.Sprintf(`Незивестнное событие: "%+v"`, event))
				}
			case state := <-client.States:
				logger.WithField("info", state).Debug("get state")
				//sendSkypeMessage(fmt.Sprintf("Текущее состояние: %s", beautyJSON(state)))
			}
		}
	}()

	log.Fatal(client.Run())
}
