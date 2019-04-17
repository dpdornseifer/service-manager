package ws

import (
	"context"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

func NewUpgrader(baseCtx context.Context, done chan struct{}) *SmUpgrader {
	return &SmUpgrader{
		baseCtx: baseCtx,
		done:    done,
	}
}

type SmUpgrader struct {
	conns     []*Conn
	connMutex sync.Mutex

	baseCtx context.Context
	done    chan struct{}
}

func (u *SmUpgrader) Upgrade(rw http.ResponseWriter, req *http.Request) (*Conn, error) {
	upgrader := &websocket.Upgrader{}
	webConn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		return nil, err
	}
	return u.addConn(webConn), nil
}

func (u *SmUpgrader) Shutdown() error {
	defer func() {
		close(u.done)
	}()

	var err error
	for _, conn := range u.conns {
		currErr := conn.conn.Close()
		if currErr != nil {
			err = currErr
		}
	}
	return err
}

func (u *SmUpgrader) addConn(c *websocket.Conn) *Conn {
	u.connMutex.Lock()
	defer u.connMutex.Unlock()
	conn := &Conn{
		conn: c,
	}
	u.conns = append(u.conns, conn)
	return conn
}
