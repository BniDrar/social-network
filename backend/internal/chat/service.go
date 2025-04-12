package chat

import (
	"bytes"
	"context"
	"log"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

func WsListing(conn *websocket.Conn, ctx context.Context)  {
	defer conn.Close()
	m := ws.NewManager()
	id := ctx.Value("userID").(uint)
	m.AddClient(id, conn)
	defer m.RemoveClient(id)
	log.Printf("Client %d connected", id)
	// Listen for messages from the client
	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		m.SendMessage(id, message)
		log.Printf("received message from Client %d: %s", id, message)
	}
}