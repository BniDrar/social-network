package chat

import (
	"bytes"
	"context"
	"errors"
	"log"

	"socialNetwork/entity"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

var Manager = ws.NewManager()

func WsListing(conn *websocket.Conn, ctx context.Context) error {
	defer conn.Close()
	val := ctx.Value(entity.ContextID)
	idInt, ok := val.(int)
	id := uint(idInt) // convert after
	if !ok {
		// handle the error: not found or wrong type
		return errors.New("user ID missing or invalid in context")
	}
	Manager.AddClient(id, conn)
	defer Manager.RemoveClient(id)
	log.Printf("Client %d connected", id)
	// Listen for messages from the client
	for {
		_, message, err := conn.ReadMessage()
		log.Printf("Client %d sent message: %s", id, message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return err
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		Manager.SendMessage(id, message)
		log.Printf("received message from Client %d: %s", id, message)
	}
}
