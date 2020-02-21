package encryption

import (
	"fmt"
	"testing"
)

var pubKey = []byte(`-----BEGIN PUBLIC KEY----- TODO -----END PUBLIC KEY-----`)

func TestRsaEncryptionService_Encode(t *testing.T) {
	rsa := NewRSAEncryptionService(pubKey)

	b, err := rsa.Encode([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("base64=%s", b)
}
