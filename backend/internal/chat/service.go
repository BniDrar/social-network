package chat

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"socialNetwork/entity"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

func WsListing(conn *websocket.Conn, ctx context.Context) error {
	defer conn.Close()
	fmt.Println("in ws listing")
	m := ws.NewManager()
	fmt.Println("new manager:", m)
	fmt.Println("before context")
	val := ctx.Value(entity.ContextID)
	idInt, ok := val.(int)
	fmt.Println("val:", val)
	fmt.Printf("val: %v, type: %T\n", val, val)
	fmt.Println("ok", ok)
	fmt.Println("id", idInt)
	id := uint(idInt) // convert after
	if !ok {
		fmt.Println("error in context")
		// handle the error: not found or wrong type
		return errors.New("user ID missing or invalid in context")
	}
	fmt.Println("after context")
	fmt.Println("id in listing", id)
	m.AddClient(id, conn)
	defer m.RemoveClient(id)
	log.Printf("Client %d connected", id)
	// Listen for messages from the client
	for {
		fmt.Println("in the loop reading messages")
		_, message, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return err
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		m.SendMessage(id, message)
		log.Printf("received message from Client %d: %s", id, message)
	}
}
