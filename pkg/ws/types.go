package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Conn struct {
	*websocket.Conn
	ID string

	Done     chan struct{}
	Shutdown chan struct{}

	work *sync.WaitGroup
}
