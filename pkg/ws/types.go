package ws

import (
	"github.com/gorilla/websocket"
)

type Conn struct {
	*websocket.Conn
	ID string

	ReadErrCh   chan error
	Shutdown    chan struct{}
	RemoteClose chan struct{}
}
