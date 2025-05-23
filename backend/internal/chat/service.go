package chat

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/utils"

	"github.com/gorilla/websocket"
)

func (c *chat) WsListing(conn *websocket.Conn, ctx context.Context) error {
	id := ctx.Value(entity.ContextID).(int)

	// Add client to hub
	c.Hub.AddClient(uint(id), conn)

	// Listen for messages from the client
	for {
		var msg entity.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.loger.Error.Printf("error: %v\n", err)
			}
			c.Hub.Unregister(uint(id))
			return err
		}
		if !utils.IsValidText(msg.Content) {
			continue
		}
		// Set sender ID
		msg.SenderID = id
		c.loger.Info.Printf("Received message: %v\n", id)

		go c.SaveMessage(ctx, msg)

		// Send message to receiver
		byteData, err := json.Marshal(msg)
		if err != nil {
			c.loger.Error.Printf("Error marshaling message: %v\n", err)
			continue
		}
		if msg.IsGroup {
			groupMembers := c.getGroupMembers(ctx, *msg.GroupID, id)
			c.Hub.SendGroupMessage(uint(*msg.GroupID), groupMembers, byteData)
			continue
		}
		c.Hub.SendPrivateMessage(uint(*msg.ReceiverID), byteData)
	}
}

func (c *chat) getUserContactsService(ctx context.Context) ([]entity.Contact, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	contacts, err := c.getContacts(ctx, userId)
	if err != nil {
		log.Println("err get contacts 0")
		return nil, http.StatusInternalServerError, err
	}
	for i, contact := range contacts {
		if c.Hub.IsOnline(uint(contact.ID)) {
			contacts[i].Online = true
		}
	}
	return contacts, http.StatusOK, nil
}

func (c *chat) getOnlineUsers(ctx context.Context) ([]entity.Contact, int, error) {
	userId := ctx.Value(entity.ContextID).(int)
	onlineIDs := c.Hub.GetOnlineUsers()
	contacts, err := c.getOnlines(ctx, onlineIDs, userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
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
