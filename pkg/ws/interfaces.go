package ws

import "net/http"

type Upgrader interface {
	Upgrade(http.ResponseWriter, *http.Request, http.Header) (*Conn, error)
	Shutdown() error
}
