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

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: false})
	logger.SetLevel(logrus.TraceLevel)

	// TODO: generate with `genKey(16)` when EAS encyption will by fixed
	key := []byte("1111111111111111")
	//key = genKey(16)

	logger.WithField("key", string(key)).Printf("generated AES key")

	pubKey := []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCYJF3lSVSZtncpyCPrFvjNiRoM
Htt4wOi7ABqZ1XOWYBoDicUyweZ2fLmxtepatHG7alnPNak441qYis7d513284Nc
31oNP8sKcwZJS//NQfgzE1H2elhfJoA2Jrrln/Z93DzrWsPD2Oddgw8YGmEHlm5U
IH6nQGf1XfuEesXycwIDAQAB
-----END PUBLIC KEY-----`)

	rsaEncryptionService := encryption.NewRSAEncryptionService(pubKey)
	aesEncryptionService := encryption.NewAESEncryptionService(key)

	client := venbest.NewClient(venbest.ClientOptions{
		ServerHost: "78.137.4.125",
		ServerPort: 19001,
		Username:   "a.shtovba",
		PPKNum:     286,
		Pwd:        "123456",
		LicenseKey: []uint{73, 10, 7, 39, 4, 50},
		Logger:     logger,
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
					sendMessage(fmt.Sprintf("Офис закрыт (%s)", event.When.Format(time.RFC1123)))
				case venbest.EventCode72:
					sendMessage(fmt.Sprintf("Офис открыт (%s)", event.When.Format(time.RFC1123)))
				case venbest.EventCode108:
					sendMessage(fmt.Sprintf("Открыта дверца ППК (%s)", event.When.Format(time.RFC1123)))
				case venbest.EventCode109:
					sendMessage(fmt.Sprintf("Закрыта дверца ППК (%s)", event.When.Format(time.RFC1123)))
				default:
					sendMessage(fmt.Sprintf(`Незивестнное событие: "%+v"`, event))
				}
			case state := <-client.States:
				logger.WithField("info", state).Debug("get state")
				sendMessage(fmt.Sprintf("Текущее состояние: %s", beautyJSON(state)))
			}
		}
	}()

	log.Fatal(client.Run())
}
