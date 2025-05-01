package chat

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"

	"socialNetwork/entity"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

var Manager = ws.NewManager()

func WsListing(conn *websocket.Conn, ctx context.Context) error {
	// defer conn.Close() i don't see why closing the connection here ?
	// and also closing it in the parent function
	// so for now i removed it from here ?

	m := Manager
	val := ctx.Value(entity.ContextID)
	idInt, ok := val.(int)
	id := uint(idInt) // convert after
	if !ok {
		// handle the error: not found or wrong type
		return errors.New("user ID missing or invalid in context")
	}
	m.AddClient(id, conn)
	defer m.RemoveClient(id)
	log.Println("after defer in wslisting:", len(m.Clients))
	// Listen for messages from the client
	for {
		log.Println("in the foor", len(m.Clients))
		_, message, err := conn.ReadMessage()
		log.Println("after reading message:", len(m.Clients))
		log.Printf("received message from Client %d: %s", id, message)
		if err != nil {
			log.Println(1, err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return err
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		// Manager.SendMessage(id, message)
		Manager.Broadcast(id, message)
		log.Println("after brodcasting:", len(m.Clients))
	}
}

func (c *chat) getUserContactsService(ctx context.Context) ([]entity.Contact, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	contacts, err := c.getContacts(ctx, userId)
	if err != nil {
		log.Println("err get contacts 0")
		return nil, http.StatusInternalServerError, err
	}
	for _, contact := range contacts {
		contact.Online = true
	}
	return contacts, http.StatusOK, nil
}
