package chat

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"socialNetwork/entity"

	"github.com/gorilla/websocket"
)

func (c *chat) WsListing(conn *websocket.Conn, ctx context.Context) error {
	val := ctx.Value(entity.ContextID)
	idInt, ok := val.(int)
	if !ok {
		return errors.New("user ID missing or invalid in context")
	}
	id := uint(idInt)

	// Add client to hub
	c.Hub.AddClient(id, conn)

	// Listen for messages from the client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return err
		}

		var msg entity.Message
		data, err := DecodeMessage(message, &msg)
		if err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		data.SenderID = idInt
		log.Printf("Received message: %+v", data)

		// Save message to database
		go c.SaveMessage(*data)

		// Send message to receiver
		byteData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			continue
		}

		c.Hub.SendPrivateMessage(uint(*data.ReceiverID), byteData)
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
