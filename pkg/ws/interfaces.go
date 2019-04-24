package ws

import (
	"context"
	"net/http"
	"sync"
)

type Upgrader interface {
	Upgrade(http.ResponseWriter, *http.Request, http.Header) (*Conn, error)
	Start(context.Context, *sync.WaitGroup)
}
