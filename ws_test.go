package util

import (
	"context"
	"fmt"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWsClient(t *testing.T) {

	wsc := &WebsocketClient{
		Url:         "ws://localhost:8000/ws",
		MessageType: websocket.TextMessage,
	}

	err := wsc.NewClient()
	if err != nil {
		t.Logf(err.Error())
		return
	}

	ctx := context.Background()

	err = wsc.Listener(ctx, WebsocketResponse{
		OnConnected: func() {
			t.Logf("connected to websocket service...")
		},
		OnMessage: func(message []byte) {
			t.Logf(fmt.Sprintf("message : %s", string(message)))
		},
		OnDisconected: func() {
			t.Logf("disconnected from websocket service...")
		},
	})
	if err != nil {
		t.Logf(err.Error())
	}

	wsc.Send([]byte("hello"))
	<-ctx.Done()

	err = wsc.CloseClient()
	if err != nil {
		t.Logf(err.Error())
	}
}
