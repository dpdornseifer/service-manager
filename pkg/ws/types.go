package ws

import (
	"github.com/gorilla/websocket"
)

type Conn struct {
	conn *websocket.Conn
}
