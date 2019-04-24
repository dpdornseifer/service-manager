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

func NewUpgrader(options *UpgraderOptions) *SmUpgrader {
	return &SmUpgrader{
		conns:       make(map[string]*Conn),
		options:     options,
		connWorkers: &sync.WaitGroup{},
	}
}

type SmUpgrader struct {
	options *UpgraderOptions

	conns       map[string]*Conn
	connMutex   sync.Mutex
	connWorkers *sync.WaitGroup

	isShutDown    bool
	shutdownMutex sync.Mutex
}

func (u *SmUpgrader) Start(baseCtx context.Context, work *sync.WaitGroup) {
	work.Add(1)
	go u.shutdown(baseCtx, work)
}

func (u *SmUpgrader) Upgrade(rw http.ResponseWriter, req *http.Request, header http.Header) (*Conn, error) {
	u.shutdownMutex.Lock()
	defer u.shutdownMutex.Unlock()
	if u.isShutDown {
		return nil, fmt.Errorf("upgrader is going to shutdown and does not accept new connections")
	}

	if header == nil {
		header = http.Header{}
	}
	header.Add("max_ping_interval", u.options.PingTimeout.String())

	upgrader := &websocket.Upgrader{}
	conn, err := upgrader.Upgrade(rw, req, header)
	if err != nil {
		return nil, err
	}
	wsConn, err := u.addConn(conn, u.connWorkers)
	if err != nil {
		return nil, err
	}
	u.setConnTimeout(wsConn)
	u.setCloseHandler(wsConn)

	u.connWorkers.Add(1)
	go u.handleConn(wsConn)

	return wsConn, nil
}

func (u *SmUpgrader) handleConn(c *Conn) {
	defer u.connWorkers.Done()
	for {
		select {
		case <-c.Done:
			u.removeConn(c.ID)
			return
		default:
		}
	}
}

func (u *SmUpgrader) shutdown(ctx context.Context, work *sync.WaitGroup) {
	<-ctx.Done()
	defer work.Done()

	u.shutdownMutex.Lock()
	u.isShutDown = true
	u.shutdownMutex.Unlock()

	for _, conn := range u.conns {
		close(conn.Shutdown)
	}
	u.connMutex.Lock()
	u.conns = nil
	u.connMutex.Unlock()

	u.connWorkers.Wait()
}

func (u *SmUpgrader) setCloseHandler(c *Conn) {
	c.SetCloseHandler(func(code int, text string) error {
		u.removeConn(c.ID)
		c.Close()
		return nil
	})
}

func (u *SmUpgrader) setConnTimeout(c *Conn) {
	c.SetReadDeadline(time.Now().Add(u.options.PingTimeout))
	// TODO: How to set write deadline, each time after write is executed?
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

func (u *SmUpgrader) addConn(c *websocket.Conn, workGroup *sync.WaitGroup) (*Conn, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	conn := &Conn{
		Conn:     c,
		ID:       uuid.String(),
		Shutdown: make(chan struct{}),
		Done:     make(chan struct{}),
		work:     workGroup,
	}

	u.connMutex.Lock()
	defer u.connMutex.Unlock()
	u.conns[conn.ID] = conn
	return conn, nil
}

func (u *SmUpgrader) removeConn(id string) {
	u.connMutex.Lock()
	defer u.connMutex.Unlock()
	delete(u.conns, id)
}
