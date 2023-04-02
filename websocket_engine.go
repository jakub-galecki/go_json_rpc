package jsonrpc

import (
	"context"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

type websocketEngine struct {
	url string
}

func newWebsocketEngine(url string) *websocketEngine {
	return &websocketEngine{url: url}
}

func (h *websocketEngine) MakeRequest(req []byte) (res Response, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c, _, err := websocket.Dial(ctx, h.url, nil)
	defer func() {
		err = c.Close(websocket.StatusInternalError, "lul")
	}()
	err = wsjson.Write(ctx, c, req)
	if err != nil {
		return Response{}, err
	}
	err = wsjson.Read(ctx, c, res)
	if err != nil {
		return Response{}, err
	}
	return
}
