package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/Luzifer/go-openssl"
)

type encryptionService struct {
	pubKeyContent []byte
	aesKeyContent []byte
}

func NewEncryptionService(pubKey []byte, AESPassword []byte) *encryptionService {
	return &encryptionService{
		aesKeyContent: AESPassword,
		pubKeyContent: pubKey,
	}
}

func (e *encryptionService) EncodeRSA(text []byte) ([]byte, error) {
	block, _ := pem.Decode(e.pubKeyContent)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	b1, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), e.aesKeyContent)
	if err != nil {
		return nil, err
	}

	return []byte(Base64Enc(b1)), nil
}

func Base64Enc(b1 []byte) string {
	s1 := base64.StdEncoding.EncodeToString(b1)
	s2 := ""
	var LEN = 76
	for len(s1) > 76 {
		s2 = s2 + s1[:LEN] + "\n"
		s1 = s1[LEN:]
	}
	s2 = s2 + s1
	return s2
}

func (e *encryptionService) EncodeAES(text []byte) ([]byte, error) {
	return openssl.New().EncryptBytes(string(e.aesKeyContent), text)
}

func (e *encryptionService) DecodeAES(text []byte) ([]byte, error) {
	return openssl.New().DecryptBytes(string(e.aesKeyContent), text)
}
