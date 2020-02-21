package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/pkg/errors"
)

type RSAEncryptionService struct {
	pubKeyContent []byte
}

func NewRSAEncryptionService(pubKey []byte) *RSAEncryptionService {
	return &RSAEncryptionService{
		pubKeyContent: pubKey,
	}
}

func (s *RSAEncryptionService) Encode(text []byte) ([]byte, error) {
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
}
