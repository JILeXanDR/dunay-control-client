package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

//var AESKey = genKey(16)
//var AESKey = []byte{49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49} // "1111111111111111"
var AESKey = bytes.NewBufferString("1111111111111111").Bytes()
var salt = []byte("Salted__")
var encService = NewEncryptionService("key.pub", salt, AESKey)

func main() {
	connection, _, err := websocket.DefaultDialer.Dial("ws://78.137.4.125:19001", nil)

	if err != nil {
		log.Fatal("dial:", err)
	}
	defer connection.Close()

	log.Printf("AESKey=%v", AESKey)
	log.Printf("AESKey=%s", AESKey)

	firstSend, err := encService.EncodeRSA(AESKey)
	if err != nil {
		log.Fatal("can't encode as RSA:", err)
	}

	firstSend = bytes.NewBufferString("A/wZaMghKZM/69rpTj9y8cAPyFIqe3IjE9wuhcUYE/DCf9xeK1GFj+iYf3FIl2epkh1dKwJusVKlKE++1LSeXW3NmlTPPxBjZdZs0Kcnu5hB+z1TVGapcHmz7wtX5oCuT6MXmlWNRionhewKCpbUP595xCbJztseowEAUSSC/fs=").Bytes()
	log.Printf("[WS] send message %v", string(firstSend))

	if err := connection.WriteMessage(websocket.TextMessage, firstSend); err != nil {
		panic("WriteMessage error -> " + err.Error())
	}

	messages := make(chan []byte)

	go func() {
		for {
			println("ReadMessage")
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("[WS] plain message: %s", message)

			println("send to channel")
			messages <- message
			println("end for...")
		}
	}()

	for {
		select {
		case message := <-messages:
			go func(message []byte) {

				defer func() {
					log.Print("Done==================================\n\n\n")
				}()

				println("channel message", string(message))

				var jsonData map[string]interface{}
				if err := json.Unmarshal(message, &jsonData); err == nil {
					log.Printf("decoded JSON response: %+v", jsonData)
					return
				}

				decoded, err := encService.DecodeAES(message)
				if err != nil {
					log.Printf("err: %v", err)
					return
				}

				if err := json.Unmarshal(decoded, &jsonData); err != nil {
					log.Printf("err: %v", err)
					return
				}

				log.Printf("json data: %+v", jsonData)

				messageType := jsonData["type"].(string)

				if messageType == "key_accepted" {

					data := []byte(`{"type":"login","user_name":"a.shtovba","ping_interval":45000,"ppks":[{"ppk_num":286,"pwd":"123456","license_key":[73,10,7,39,4,50]}]}`)

					encoded, err := encService.EncodeAES(data)
					if err != nil {
						log.Printf("err: %v", err)
						return
					}

					encoded = []byte(`U2FsdGVkX19U+bdUkih6g4ejjpnYYGH2cw0c5LFJ98dczaZv9+KfEzCLPviqQWRFxQR59UP8rHI7zOrZiqs6+ijAHyF0Z34SHquc2dTTXyCwQ9jvyX2tmKJSHWVV+KsR23goCYOyJXA1/kPr9KrV0o5AQLbTza55465BUdMENOP7HiDiVWQRjozDxCdiLHx0MP+Jx1DN21WTkyV7RDVHtw==`)

					log.Printf("[WS] send login data %s", encoded)

					if err := connection.WriteMessage(websocket.TextMessage, encoded); err != nil {
						log.Printf("err: %v", err)
						return
					}

					log.Println("message sent")

				} else if messageType == "error" {
					log.Printf("got error from WS server: %v", err)
				} else {
					println("can't handle json response")
				}
			}(message)
		default:
			continue
		}
	}
}
