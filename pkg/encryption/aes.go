package encryption

import (
	"github.com/Luzifer/go-openssl"
)

type aesEncryptionService struct {
	key []byte
}

func NewAESEncryptionService(key []byte) *aesEncryptionService {
	return &aesEncryptionService{
		key: key,
	}
}

func (e *aesEncryptionService) Encode(text []byte) ([]byte, error) {
	return openssl.New().EncryptBytes(string(e.key), text)
}

func (e *aesEncryptionService) Decode(text []byte) ([]byte, error) {
	return openssl.New().DecryptBytes(string(e.key), text)
}
