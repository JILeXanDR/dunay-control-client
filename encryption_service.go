package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	openssl "github.com/Luzifer/go-openssl"
	"io/ioutil"
	"log"
	"os"
)

type encryptionService struct {
	RSAKey  *rsa.PublicKey
	RSASalt []byte
	AESKey  []byte
}

func NewEncryptionService(pubRSAKeyPath string, RSASalt []byte, AESPassword []byte) *encryptionService {
	s := encryptionService{
		RSASalt: RSASalt,
		AESKey:  AESPassword,
	}

	s.RSAKey = s.getPubKey(pubRSAKeyPath)

	return &s
}

func (e *encryptionService) EncodeRSA(data []byte) ([]byte, error) {
	rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(sha1.New(), rng, e.RSAKey, data, e.RSASalt)
	if err != nil {
		return nil, err
	}

	return bytes.NewBufferString(base64.StdEncoding.EncodeToString(ciphertext)).Bytes(), nil
}

func (e *encryptionService) getPubKey(path string) *rsa.PublicKey {
	file, _ := os.Open(path)
	b, _ := ioutil.ReadAll(file)
	test2048Key, _ := e.bytesToPublicKey(b)
	return test2048Key
}

func (e *encryptionService) EncodeAES(text []byte) ([]byte, error) {
	o := openssl.New()

	enc, err := o.EncryptBytes(string(e.AESKey), text)
	if err != nil {
		return nil, err
	}

	return enc, nil
}

func (e *encryptionService) DecodeAES(ciphertext []byte) ([]byte, error) {
	o := openssl.New()

	dec, err := o.DecryptBytes(string(e.AESKey), ciphertext)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

	return dec, nil
}

func (e *encryptionService) bytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not PublicKey")
	}
	return key, nil
}

func genKey(length uint) []byte {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	return key
	//return bytes.NewBufferString(base64.StdEncoding.EncodeToString(key)).Bytes()
}
