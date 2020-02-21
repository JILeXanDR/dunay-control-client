package venbest

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

type ws struct {
	addr           *url.URL
	processMessage chan []byte
	// conn represents WS connection with server
	conn   *websocket.Conn
	logger *logrus.Logger
}

func (w *ws) connect() (func(), error) {
	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
	}
	conn, _, err := dialer.Dial(w.addr.String(), nil)
	if err != nil {
		return nil, err
	}

	w.conn = conn

	closeFunc := func() {
		if err := conn.Close(); err != nil {
			w.logger.WithError(err).Error("failed to close WS connection")
		}
	}

	return closeFunc, nil
}

func (w *ws) readMessages() {
	defer func() {
		w.logger.Warn("stop reading of messages")
	}()
	for {
		_, message, err := w.conn.ReadMessage()
		if err != nil {
			w.logger.WithError(err).Printf("can't read message")
			break
		}

		w.logger.WithField("data", message).Printf("Received plain message")

		w.processMessage <- message
	}
}

func (w *ws) send(data []byte) error {
	w.logger.WithField("data", data).Printf("send message")
	return w.conn.WriteMessage(websocket.TextMessage, data)
}
