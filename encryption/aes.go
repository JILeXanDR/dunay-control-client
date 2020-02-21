package encryption

import (
	"github.com/Luzifer/go-openssl"
)

type AESEncryptionService struct {
	key []byte
}

func NewAESEncryptionService(key []byte) *AESEncryptionService {
	return &AESEncryptionService{
		key: key,
	}
}

func (e *AESEncryptionService) Encode(text []byte) ([]byte, error) {
	return openssl.New().EncryptBytes(string(e.key), text)
}

func (e *AESEncryptionService) Decode(text []byte) ([]byte, error) {
	return openssl.New().DecryptBytes(string(e.key), text)
}
