package util

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	Url         string
	MessageType int
	conn        *websocket.Conn
}

type WebsocketResponse struct {
	OnConnected   func()
	OnMessage     func(message []byte)
	OnDisconected func()
}

func (c *WebsocketClient) NewClient() error {

	dialer := websocket.Dialer{
		Subprotocols:    []string{},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	header := http.Header{"Accept-Encoding": []string{"gzip"}}
	conn, _, err := dialer.Dial(c.Url, header)
	if err != nil {
		return err
	}

	c.conn = conn

	return nil
}

func (c *WebsocketClient) CloseClient() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}

	c.conn = nil

	return nil
}

func (c *WebsocketClient) Listener(ctx context.Context, r WebsocketResponse) error {

	if c.conn == nil {
		return errors.New("Instance is nil!")
	}

	r.OnConnected()
	defer r.OnDisconected()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:

			typeM, message, err := c.conn.ReadMessage()
			if err != nil {
				return err
			}

			if typeM != c.MessageType {
				return errors.New("not expected message type!")
			}

			r.OnMessage(message)

		}
	}
}

func (c *WebsocketClient) Send(message []byte) error {
	if c.conn == nil {
		return errors.New("Instance is nil!")
	}

	err := c.conn.WriteMessage(c.MessageType, message)
	if err != nil {
		return err
	}

	return nil

}
