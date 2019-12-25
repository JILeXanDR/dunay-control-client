package encryption

import (
	"fmt"
	"testing"
)

var pubKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCYJF3lSVSZtncpyCPrFvjNiRoM
Htt4wOi7ABqZ1XOWYBoDicUyweZ2fLmxtepatHG7alnPNak441qYis7d513284Nc
31oNP8sKcwZJS//NQfgzE1H2elhfJoA2Jrrln/Z93DzrWsPD2Oddgw8YGmEHlm5U
IH6nQGf1XfuEesXycwIDAQAB
-----END PUBLIC KEY-----`)

func TestRsaEncryptionService_Encode(t *testing.T) {
	rsa := NewRSAEncryptionService(pubKey)

	b, err := rsa.Encode([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("base64=%s", b)
}