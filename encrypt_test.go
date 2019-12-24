package main

import (
	"fmt"
	"testing"
)

var AESKey = []byte("1111111111111111")
var pubKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCYJF3lSVSZtncpyCPrFvjNiRoM
Htt4wOi7ABqZ1XOWYBoDicUyweZ2fLmxtepatHG7alnPNak441qYis7d513284Nc
31oNP8sKcwZJS//NQfgzE1H2elhfJoA2Jrrln/Z93DzrWsPD2Oddgw8YGmEHlm5U
IH6nQGf1XfuEesXycwIDAQAB
-----END PUBLIC KEY-----`)

func TestEncryptionService_EncodeAES(t *testing.T) {
	testService := NewEncryptionService(pubKey, AESKey)

	message := []byte(`xxx`)

	{
		encoded, err := testService.EncodeAES(message)
		if err != nil {
			t.Fatal(err)
		}
		println("encoded 1", string(encoded))
	}

	{
		encoded, err := testService.EncodeAES(message)
		if err != nil {
			t.Fatal(err)
		}
		println("encoded 2", string(encoded))
	}
}

func TestEncryptionService_DecodeAES(t *testing.T) {
	testService := NewEncryptionService(pubKey, AESKey)

	word := []byte(`{"type":"login","user_name":"a.shtovba","ping_interval":45000,"ppks":[{"ppk_num":286,"pwd":"123456","license_key":[73,10,7,39,4,50]}]}`)

	encoded, err := testService.EncodeAES(word)
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := testService.DecodeAES(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if string(decoded) != string(word) {
		t.Fatal("bad result")
	}
	println("decoded", string(decoded))
}

func TestEncryptionService_DecodeAES1(t *testing.T) {
	testService := NewEncryptionService(pubKey, []byte("FPLLNGZIEYOH43E0"))

	encoded := []byte(`U2FsdGVkX1/yP0XzjbIecUF2t2IwnzZk+9Zd/QNqtJQMewgb1Qp+gLP8CvciVx4jVDfym0ZVuu2fAbtL6MK0mgiaRn1wvo1s5aiZdR9Zdi6jYtceo7Gu2YFmjBV7CSaypYE7moARipfzz0TX0B3ktxibBzvhFW0PypruevjT32EjxUrtjEmzrzvjzrfVhFe6lqMCQ7Zhzow/go3FkpzsUw==`)

	{
		decoded1, err := testService.DecodeAES(encoded)
		if err != nil {
			t.Fatal(err)
		}
		println("decoded 1", string(decoded1))
	}

	{
		decoded2, err := testService.DecodeAES(encoded)
		if err != nil {
			t.Fatal(err)
		}
		println("decoded 2", string(decoded2))
	}
}

func TestEncryptionService_EncryptRSA(t *testing.T) {
	testService := NewEncryptionService(pubKey, AESKey)

	b, err := testService.EncodeRSA(AESKey)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("base64=%s", b)
}
