package ws

type SmHandler struct{}

func (h *SmHandler) Run(wsconn *Conn) error {
	wsconn.conn.WriteJSON("hello")
	return nil
}
