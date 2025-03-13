package websocket

type Hub struct {
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) Insert() {

}