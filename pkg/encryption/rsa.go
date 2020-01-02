package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/pkg/errors"
)

type rsaEncryptionService struct {
	pubKeyContent []byte
}

func NewRSAEncryptionService(pubKey []byte) *rsaEncryptionService {
	return &rsaEncryptionService{
		pubKeyContent: pubKey,
	}
}

func (s *rsaEncryptionService) Encode(text []byte) ([]byte, error) {
	var block *pem.Block
	if val, err := pem.Decode(s.pubKeyContent); len(err) != 0 {
		return nil, errors.Wrap(errors.New(string(err)), "failed to decode RSA public key")
	} else if val == nil {
		return nil, errors.New("public key error")
	} else {
		block = val
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ParsePKIXPublicKey")
	}

	b1, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), text)
	if err != nil {
		return nil, errors.Wrap(err, "failed to EncryptPKCS1v15")
	}

	return []byte(base64.StdEncoding.EncodeToString(b1)), nil
	//return base64Encode(b1), nil
}

/*func base64Encode(data []byte) []byte {
	res := base64.StdEncoding.EncodeToString(data)
	return []byte(res)
	//s2 := ""
	//var LEN = 76
	//for len(s1) > 76 {
	//	s2 = s2 + s1[:LEN] + "\n"
	//	s1 = s1[LEN:]
	//}
	//s2 = s2 + s1
	//return s2
}
*/
