package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"strings"
)

func genKey(length int) []byte {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return []byte(b.String())
}

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logger.SetLevel(logrus.TraceLevel)

	key := []byte("1111111111111111")
	//key = genKey(16)

	logger.WithField("key", string(key)).Printf("generated AES key")

	pubKey := []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCYJF3lSVSZtncpyCPrFvjNiRoM
Htt4wOi7ABqZ1XOWYBoDicUyweZ2fLmxtepatHG7alnPNak441qYis7d513284Nc
31oNP8sKcwZJS//NQfgzE1H2elhfJoA2Jrrln/Z93DzrWsPD2Oddgw8YGmEHlm5U
IH6nQGf1XfuEesXycwIDAQAB
-----END PUBLIC KEY-----`)

	encryptor := NewEncryptionService(pubKey, key)

	client := NewVinbestClient(vinbestClientOptions{
		ServerHost:       "78.137.4.125",
		ServerPort:       19001,
		AESEncryptionKey: key,
		Username:         "a.shtovba",
		PPKNum:           286,
		Pwd:              "123456",
		LicenseKey:       []int{73, 10, 7, 39, 4, 50},
		Logger:           logger,
		MessageEncoder: func(message []byte) ([]byte, error) {
			return encryptor.EncodeAES(message)
		},
		MessageDecoder: func(message []byte) ([]byte, error) {
			return encryptor.DecodeAES(message)
		},
		KeyPreparer: func(key []byte) ([]byte, error) {
			return encryptor.EncodeRSA(key)
		},
	})

	log.Fatal(client.Run())
}
