package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/gorilla/websocket"
	"io"
	"log"
)

var secretKey = []byte("6368616e676520746869732070617373776f726420746f206120736563726574")

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://78.137.4.125:19001", nil)

	// TODO: send encrypted secret key

	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	if err := c.WriteMessage(websocket.TextMessage, bytes.NewBufferString(Encrypt(string(secretKey))).Bytes()); err != nil {
		panic("WriteMessage error -> " + err.Error())
	}

	for {
		println("ReadMessage")
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		log.Printf("response: %s", message)

		Decrypt(string(message))
	}

	println("exit")
}

func Encrypt(data string) string {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString(string(secretKey))
	plaintext := []byte(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return string(ciphertext)
}

func Decrypt(data string) string {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString(string(secretKey))
	ciphertext, _ := hex.DecodeString(data)
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
