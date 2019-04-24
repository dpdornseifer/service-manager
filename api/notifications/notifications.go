package notifications

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	"github.com/Peripli/service-manager/notifications"

	"github.com/Peripli/service-manager/pkg/query"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"

	"github.com/Peripli/service-manager/pkg/log"

	"github.com/Peripli/service-manager/pkg/web"
	"github.com/Peripli/service-manager/pkg/ws"
)

func (c *Controller) handleWS(req *web.Request) (*web.Response, error) {
	user, _ := web.UserFromContext(req.Context())
	notificationQueue, lastKnownRevision, err := c.notificator.RegisterConsumer(user)
	if err != nil {
		return nil, fmt.Errorf("could not register notification consumer: %v", err)
	}

	revisionKnownToProxy := 0
	revisionKnownToProxyStr := req.URL.Query().Get("last_known_revision")
	if revisionKnownToProxyStr != "" {
		revisionKnownToProxy, err = strconv.Atoi(revisionKnownToProxyStr)
		if err != nil {
			c.unregisterConsumer(notificationQueue)

			log.C(req.Context()).Errorf("could not convert string to number: %v", err)
			return nil, &util.HTTPError{
				StatusCode:  http.StatusBadRequest,
				Description: fmt.Sprintf("invalid last_known_revision query parameter"),
				ErrorType:   "BadRequest",
			}
		}
	}

	listQuery1 := query.ByField(query.GreaterThanOperator, "revision", strconv.Itoa(revisionKnownToProxy))
	listQuery2 := query.ByField(query.LessThanOperator, "revision", strconv.FormatInt(lastKnownRevision, 10))
	notificationsList, err := c.repository.List(req.Context(), types.NotificationType, listQuery1, listQuery2)
	if err != nil {
		// TODO: Wrap err
		return nil, err
	}

	rw := req.HijackResponseWriter()
	conn, err := c.wsUpgrader.Upgrade(rw, req.Request, http.Header{
		"last_known_revision": []string{strconv.FormatInt(lastKnownRevision, 10)},
	})
	if err != nil {
		c.unregisterConsumer(notificationQueue)
		return nil, err
	}

	go c.writeLoop(conn, notificationsList, notificationQueue)
	go c.readLoop(conn)

	return &web.Response{}, nil
}

func (c *Controller) readLoop(conn *ws.Conn) {
	defer func() {
		close(conn.Done)
		conn.Close()
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.D().Errorf("ws: could not read: %v", err)
			return
		}
	}
}

func (c *Controller) writeLoop(conn *ws.Conn, notificationsList types.ObjectList, q notifications.NotificationQueue) {
	defer c.unregisterConsumer(q)
	defer func() {
		close(conn.Done)
		conn.Close()
	}()

	for i := 0; i < notificationsList.Len(); i++ {
		notification := (notificationsList.ItemAt(i)).(*types.Notification)
		select {
		case <-conn.Shutdown:
			c.sendWsClose(conn)
			return
		default:
		}
		if err := conn.WriteJSON(notification); err != nil {
			log.D().Errorf("ws: could not write: %v", err)
			return
		}
	}

	notificationChannel, err := q.Channel()
	if err != nil {
		log.D().Errorf("Could not aquire notification channel: %v", err)
		return
	}

	for {
		select {
		case <-conn.Shutdown:
			c.sendWsClose(conn)
			return
		case notification, ok := <-notificationChannel:
			if !ok {
				c.sendWsClose(conn)
				return
			}
			if err := conn.WriteJSON(notification); err != nil {
				log.D().Errorf("ws: could not write: %v", err)
				return
			}
		}
	}
}

func (c *Controller) sendWsClose(conn *ws.Conn) {
	// TODO: Timeout?
	if err := conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, ""), time.Time{}); err != nil {
		log.D().Errorf("Could not send close message: %v", err)
	}
}

func (c *Controller) unregisterConsumer(q notifications.NotificationQueue) {
	if unregErr := c.notificator.UnregisterConsumer(q); unregErr != nil {
		log.D().Errorf("Could not unregister notification consumer: %v", unregErr)
	}
}
