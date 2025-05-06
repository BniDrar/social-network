package chat

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"socialNetwork/entity"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

var Manager = ws.NewManager()

func (c *chat) WsListing(conn *websocket.Conn, ctx context.Context) error {
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
	// Listen for messages from the client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		var msg entity.Message
		data, err := DecodeMessage(message, &msg)
		data.SenderID = idInt
		log.Println("recieved package:", data)
		if err != nil {
			log.Println(msg)
		}

		if err != nil {
			log.Println(1, err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return err
		}
		// message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		byteData, _ := json.Marshal(data)
		go c.SaveMessage(*data)
		go Manager.SendMessage(uint(*data.ReceiverID), byteData)
		log.Println("message sent:", len(m.Clients))
	}
}

func (c *chat) getUserContactsService(ctx context.Context) ([]entity.Contact, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	contacts, err := c.getContacts(ctx, userId)
	if err != nil {
		log.Println("err get contacts 0")
		return nil, http.StatusInternalServerError, err
	}
	for i := range contacts {
		contacts[i].Online = true
	}
	return contacts, http.StatusOK, nil
}

func (c *chat) getOnlineUsers(ctx context.Context) ([]entity.Contact, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	contacts, err := c.getOnlines(ctx, userId)
	if err != nil {
		log.Println("err get contacts 0")
		return nil, http.StatusInternalServerError, err
	}
	for i := range contacts {
		contacts[i].Online = true
	}
	return contacts, http.StatusOK, nil
}

func (c *chat) getMessagesService(ctx context.Context, cursor entity.Cursor) (int, []entity.Message, error) {
	userId := ctx.Value(entity.ContextID).(int)
	if (cursor.Time == nil && cursor.LastId != nil) || (cursor.Time != nil && cursor.LastId == nil) {
		return http.StatusBadRequest, nil, errors.New("invalid request body")
	}
	messages, err := c.getMessagesRepo(ctx, cursor, userId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, messages, nil
}
