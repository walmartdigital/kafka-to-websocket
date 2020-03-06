package server

import (
	"fmt"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type conn struct {
	id           string
	ws           *websocket.Conn
	channel      chan []byte
	closeHandler func()
}

func createConn(ws *websocket.Conn) *conn {
	channel := make(chan []byte)
	conn := &conn{uuid.NewV4().String(), ws, channel, func() {}}
	go conn.readWebsocketMessageLoop()
	go conn.readChannelMessageLoop()

	return conn
}

func (c *conn) readChannelMessageLoop() {
	for {
		select {
		case message := <-c.channel:
			fmt.Println(string(message))
			err := c.ws.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				break
			}
		}
	}
}

func (c *conn) readWebsocketMessageLoop() {
	for {
		kind, value, err := c.ws.ReadMessage()
		c.channel <- value
		if kind == websocket.CloseMessage || err != nil {
			c.closeHandler()
			return
		}
	}
}

func (c *conn) setCloseHandler(f func()) {
	c.closeHandler = f
}
