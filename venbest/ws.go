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

func newWS(addr *url.URL, logger *logrus.Logger) *ws {
	return &ws{
		addr:           addr,
		processMessage: make(chan []byte),
		logger:         logger.WithField("module", "ws").Logger,
	}
}

func (w *ws) connect() (func(), error) {
	w.logger.Debugf(`connecting to %s`, w.addr.String())
	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
	}
	conn, _, err := dialer.Dial(w.addr.String(), nil)
	if err != nil {
		return nil, err
	}

	w.conn = conn

	closeConn := func() {
		w.logger.Debug("closing connection with server...")
		if err := conn.Close(); err != nil {
			w.logger.WithError(err).Error("failed to close connection")
		}
	}

	return closeConn, nil
}

func (w *ws) readMessages() {
	var lastErr error
	defer func() {
		w.logger.WithError(lastErr).Warn("stop reading of messages")
	}()
	for {
		w.logger.Debug("read message...")
		_, message, err := w.conn.ReadMessage()
		if err != nil {
			lastErr = err
			w.logger.WithError(err).Printf("can't read message")
			break
		}

		w.logger.WithField("data", string(message)).Printf("received plain message")

		w.logger.Debugf("before send: %s", message)
		w.processMessage <- message
		w.logger.Debug("after send")
	}
}

func (w *ws) send(data []byte) error {
	w.logger.WithField("data", string(data)).Printf("send message")
	return w.conn.WriteMessage(websocket.TextMessage, data)
}
