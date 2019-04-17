package notifications

import (
	"time"

	"github.com/Peripli/service-manager/pkg/web"
)

func (c *Controller) handleTest(req *web.Request) (*web.Response, error) {
	time.Sleep(time.Second * 10)
	return &web.Response{
		StatusCode: 200,
	}, nil
}

func (c *Controller) handleWS(req *web.Request) (*web.Response, error) {
	rw := req.HijackResponseWriter()
	wsConn, err := c.wsUpgrader.Upgrade(rw, req.Request)
	if err != nil {
		return nil, err
	}

	if err := c.wsHandler.Run(wsConn); err != nil {
		return nil, err
	}

	// upgrader := &websocket.Upgrader{}
	// webConn, err := upgrader.Upgrade(rw, req.Request, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// reqCtx := req.Context()

	// go func() {
	// 	<-c.baseCtx.Done()
	// 	log.C(reqCtx).Info("ws: interrupt received")
	// 	time.Sleep(time.Second * 3)

	// 	log.C(reqCtx).Info("websocket conn closed")
	// 	err := webConn.Close()
	// 	if err != nil {
	// 		log.C(reqCtx).Errorf("websocket: %s", err)
	// 	}
	// 	close(c.wsCh)
	// }()

	// webConn.WriteJSON("Hello from the other siiiiiide")

	// go func() {
	// 	for {
	// 		mt, message, err := webConn.ReadMessage()
	// 		if err != nil {
	// 			fmt.Println("read:", err)
	// 			break
	// 		}
	// 		log.C(req.Context()).Info(">>>>>>>>>>>>>>>>here")
	// 		fmt.Printf("recv: %s\n", message)
	// 		err = webConn.WriteMessage(mt, message)
	// 		if err != nil {
	// 			fmt.Println("write:", err)
	// 			break
	// 		}
	// 		select {
	// 		case <-c.wsCh:
	// 			break
	// 		case <-time.After(time.Second):
	// 		}
	// 	}
	// }()

	return &web.Response{}, nil
}
