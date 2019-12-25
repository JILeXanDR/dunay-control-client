package main

import (
	"dunay-control-client/pkg/encryption"
	"dunay-control-client/pkg/venbest"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"net/url"
)

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logger.SetLevel(logrus.TraceLevel)

	// TODO: generate with `genKey(16)` when EAS encyption will by fixes
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
		MessageEncoder: func(message []byte) ([]byte, error) {
			b, err := aesEncryptionService.Encode(message)
			if err != nil {
				return nil, err
			}

			return []byte(base64.StdEncoding.EncodeToString(b)), nil
		},
		MessageDecoder: func(message []byte) ([]byte, error) {
			return aesEncryptionService.Decode(message)
		},
		KeyPreparer: func() ([]byte, error) {
			return rsaEncryptionService.Encode(key)
		},
		StateHandler: func(payload map[string]interface{}) {
			logger.WithField("payload", payload).Debug("state handler")
			sendMessage(fmt.Sprintf("state:\n\n%s", beautyJSON(payload)))
		},
		EventsHandler: func(payload map[string]interface{}) {
			logger.WithField("payload", payload).Debug("events handler")
			sendMessage(fmt.Sprintf("event:\n\n%s", beautyJSON(payload)))
		},
	})

	log.Fatal(client.Run())
}

func beautyJSON(v interface{}) []byte {
	j, _ := json.MarshalIndent(v, "", "    ")
	return j
}

func sendMessage(text string) {
	http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%v&text=%s", "846079612:AAEAH62nzgQWY51hEoRsupx9OgdDIPAsMl8", 253637452, url.QueryEscape(text)))
}
