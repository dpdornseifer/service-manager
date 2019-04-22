package ws

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gofrs/uuid"

	"github.com/gorilla/websocket"
)

type UpgraderOptions struct {
	PingTimeout  time.Duration `mapstructure:"ping_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func NewUpgrader(baseCtx context.Context, done chan struct{}, options *UpgraderOptions) *SmUpgrader {
	return &SmUpgrader{
		conns:   make(map[string]*Conn),
		baseCtx: baseCtx,
		done:    done,
		options: options,
	}
}

type SmUpgrader struct {
	options *UpgraderOptions

	conns     map[string]*Conn
	connMutex sync.Mutex

	baseCtx       context.Context
	done          chan struct{}
	isShutDown    bool
	shutdownMutex sync.Mutex
}

func (u *SmUpgrader) Upgrade(rw http.ResponseWriter, req *http.Request, header http.Header) (*Conn, error) {
	u.shutdownMutex.Lock()
	defer u.shutdownMutex.Unlock()
	if u.isShutDown {
		return nil, fmt.Errorf("upgrader is going to shutdown and does not accept new connections")
	}

	upgrader := &websocket.Upgrader{}
	conn, err := upgrader.Upgrade(rw, req, header)
	if err != nil {
		return nil, err
	}
	wsConn, err := u.addConn(conn)
	if err != nil {
		return nil, err
	}
	u.setConnTimeout(wsConn)
	u.setCloseHandler(wsConn)
	go u.handleConn(wsConn)

	return wsConn, nil
}

func (u *SmUpgrader) handleConn(c *Conn) {
	for {
		select {
		case <-c.ReadErrCh:
			u.RemoveConn(c.ID)
			return
		default:
		}
	}
}

func (u *SmUpgrader) RemoveConn(id string) {
	u.connMutex.Lock()
	defer u.connMutex.Unlock()
	delete(u.conns, id)
}

func (u *SmUpgrader) Shutdown() error {
	defer func() {
		close(u.done)
	}()
	u.shutdownMutex.Lock()
	u.isShutDown = true
	u.shutdownMutex.Unlock()

	var err error
	for _, conn := range u.conns {
		close(conn.Shutdown)
	}
	u.conns = nil

	return err
}

func (u *SmUpgrader) setCloseHandler(c *Conn) {
	c.SetCloseHandler(func(code int, text string) error {
		u.RemoveConn(c.ID)
		close(c.RemoteClose)
		return nil
	})
}

func (u *SmUpgrader) setConnTimeout(c *Conn) {
	c.SetReadDeadline(time.Now().Add(u.options.PingTimeout))
	// TODO: How to set write deadline?
	// c.SetWriteDeadline(time.Now().Add(u.options.WriteTimeout))

	c.SetPingHandler(func(message string) error {
		c.SetReadDeadline(time.Now().Add(u.options.PingTimeout))

		// TODO: Deadline?
		err := c.WriteControl(websocket.PongMessage, []byte(message), time.Time{})
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	})
}

func (u *SmUpgrader) addConn(c *websocket.Conn) (*Conn, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	conn := &Conn{
		Conn:        c,
		ID:          uuid.String(),
		Shutdown:    make(chan struct{}),
		RemoteClose: make(chan struct{}),
		ReadErrCh:   make(chan error),
	}

	u.connMutex.Lock()
	defer u.connMutex.Unlock()
	u.conns[conn.ID] = conn
	return conn, nil
}
