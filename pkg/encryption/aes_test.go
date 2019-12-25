package encryption

import (
	"testing"
)

var aes = NewAESEncryptionService([]byte("1111111111111111"))

func TestAesEncryptionService_Encode(t *testing.T) {
	{
		encoded, err := aes.Encode([]byte(`{"type":"login","user_name":"a.shtovba","ping_interval":45000,"ppks":[{"ppk_num":286,"pwd":"123456","license_key":[73,10,7,39,4,50]}]}`))
		if err != nil {
			t.Fatal(err)
		}
		println("encoded 1", string(encoded))
	}

	{
		encoded, err := aes.Encode([]byte(`hello!`))
		if err != nil {
			t.Fatal(err)
		}
		println("encoded 2", string(encoded))
	}

	{
		encoded, err := aes.Encode([]byte(`hello!`))
		if err != nil {
			t.Fatal(err)
		}
		println("encoded 3", string(encoded))
	}
}

func TestAesEncryptionService_Decode(t *testing.T) {
	plainJSON := []byte(`{"type":"login","user_name":"a.shtovba","ping_interval":45000,"ppks":[{"ppk_num":286,"pwd":"123456","license_key":[73,10,7,39,4,50]}]}`)

	encoded, err := aes.Encode(plainJSON)
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := aes.Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if string(decoded) != string(plainJSON) {
		t.Fatal("bad result")
	}
	println("decoded", string(decoded))
}

func TestAesEncryptionService_Decode2(t *testing.T) {
	encoded := []byte(`U2FsdGVkX1/yP0XzjbIecUF2t2IwnzZk+9Zd/QNqtJQMewgb1Qp+gLP8CvciVx4jVDfym0ZVuu2fAbtL6MK0mgiaRn1wvo1s5aiZdR9Zdi6jYtceo7Gu2YFmjBV7CSaypYE7moARipfzz0TX0B3ktxibBzvhFW0PypruevjT32EjxUrtjEmzrzvjzrfVhFe6lqMCQ7Zhzow/go3FkpzsUw==`)

	{
		decoded1, err := aes.Decode(encoded)
		if err != nil {
			t.Fatal(err)
		}
		println("decoded 1", string(decoded1))
	}

	{
		decoded2, err := aes.Decode(encoded)
		if err != nil {
			t.Fatal(err)
		}
		println("decoded 2", string(decoded2))
	}
}
