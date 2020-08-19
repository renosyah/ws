# Websocket Library for internal use only

- how to use : 

```

import "github.com/renosyah/ws"

```

- Websocket Server 

```

	ws := &ws.WebsocketServer{}
	ws.Init()

	http.Handle("/ws", ws)
	http.ListenAndServe(":8000", nil)

```


- Websocket Client

```
    ctx := context.Background()


	wsc := &ws.WebsocketClient{
		Url : "ws://localhost:8000/ws",
		MessageType : ws.MessageTypeText,
	}


    err := wsc.NewClient()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = wsc.Listener(ctx, ws.WebsocketResponse{
		OnConnected: func() {
			fmt.Println("connected to websocket service...")
		},
		OnMessage: func(message []byte) {
			fmt.Println(fmt.Sprintf("message : %s", string(message)))
		},
		OnDisconected: func() {
			fmt.Println("disconnected from websocket service...")
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	<-ctx.Done()

	err = wsc.CloseClient()
	if err != nil {
		fmt.Println(err.Error())
	}

```