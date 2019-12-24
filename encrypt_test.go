package main

import (
	"bytes"
	"fmt"
	"testing"
)

var testService = NewEncryptionService("key.pub", salt, AESKey)

func TestEncryptionService_EncodeAES(t *testing.T) {
	encoded, err := testService.EncodeAES([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	println("res", string(encoded))
}

func TestEncryptionService_DecodeAES(t *testing.T) {
	word := []byte(`{"type": "message"}`)

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
	encoded := bytes.NewBufferString(`U2FsdGVkX19IkWYtuAoa48NcFCvUjNnPwkB/fCwucj3hGaEANAXzudmTW26ZGT/shLJw+/yInFteGuP1xkt/ijQRm8ip0qJQfjRAvstQYCw=`).Bytes()

	decoded, err := testService.DecodeAES(encoded)
	if err != nil {
		t.Fatal(err)
	}

	println("decoded", string(decoded))
}

func TestEncryptionService_EncryptRSA(t *testing.T) {
	b, err := testService.EncodeRSA(AESKey)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("base64=%s", b)
}
