package api

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type WSEvents struct {
	conn *websocket.Conn
}

func CreateTxSubscription(url string) (*WSEvents, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &WSEvents{
		conn: conn,
	}, nil
}

func (we *WSEvents) ReadCycle() {
	for {
		mtype, msg, err := we.conn.ReadMessage()
		fmt.Printf("WS EVENT type=%d, data=%s, err=%v\n", mtype, string(msg), err)
	}
}
