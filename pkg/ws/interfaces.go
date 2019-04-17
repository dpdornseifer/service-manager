package ws

import "net/http"

type Upgrader interface {
	Upgrade(http.ResponseWriter, *http.Request) (*Conn, error)
	Shutdown() error
}

type Handler interface {
	Run(*Conn) error
}
