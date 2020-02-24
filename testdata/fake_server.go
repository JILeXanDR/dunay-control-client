package testdata

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func NewFakeServer() *httptest.Server {
	upgrader := websocket.Upgrader{}

	debug := func(format string, v ...interface{}) {
		log.Printf("[FAKE_SERVER] %s", fmt.Sprintf(format, v...))
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		debug("incoming connection...")
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		debug("connection is established")

		sendTextMessage := func(text string) error {
			debug("send message: %s", text)
			return c.WriteMessage(websocket.TextMessage, []byte(text))
		}

		go func() {
			for {
				time.Sleep(3 * time.Second)
				//sendTextMessage("pong")
			}
		}()

		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}

			text := string(message)

			debug("got message: %s", text)

			if text == "ping" {
				sendTextMessage(`pong`)
			} else if text == `keyhere` {
				sendTextMessage(`{"type":"key_accepted"}`)
			} else {
				sendTextMessage(`{"type":"error","message":"bad incoming message"}`)
			}
		}
	}))
}
