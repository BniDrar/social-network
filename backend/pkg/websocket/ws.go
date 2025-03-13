package websocket

import (
	"bytes"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type client struct {
	conn *websocket.Conn
	send chan []byte
}

type WsManager struct {
	clients map[uint]*client
	mu      sync.Mutex
	hub     *Hub
}

func NewWsManager(hub *Hub) *WsManager {
	return &WsManager{
		clients: make(map[uint]*client),
		hub:     hub,
	}
}

func (m *WsManager) AddClient(userID uint, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c := &client{
		conn: conn,
		send: make(chan []byte, 1024),
	}
	m.clients[userID] = c
	log.Printf("Client %d added", userID)
	go m.readMessage(userID)
	go m.writeMessage(userID)
}

func (m *WsManager) RemoveClient(userID uint) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if c, ok := m.clients[userID]; ok {
		close(c.send)
		delete(m.clients, userID)
		log.Printf("Client %d removed", userID)
	}
}

func (m *WsManager) readMessage(userID uint) {
	c := m.clients[userID]
	defer m.RemoveClient(userID)

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from client %d: %v", userID, err)
			break
		}
		log.Printf("Received message from client %d: %s", userID, string(message))

		// Handle the message (e.g., broadcast to other clients or process it)
		m.Broadcast(userID, message)
	}
}

func (m *WsManager) writeMessage(userID uint) {
	c := m.clients[userID]
	defer m.RemoveClient(userID)

	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing message to client %d: %v", userID, err)
			break
		}
	}
}

func (m *WsManager) SendMessage(userID uint, message []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if c, ok := m.clients[userID]; ok {
		select {
		case c.send <- message:
		default:
			log.Printf("Client %d send buffer is full", userID)
		}
	}
}

func (m *WsManager) Broadcast(userID uint, message []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for id, c := range m.clients {
		if id != userID {
			select {
			case c.send <- message:
			default:
				log.Printf("Client %d send buffer is full", id)
			}
		}
	}
}

func (m *WsManager) GetClients() map[uint]*client {
	return m.clients
}

func (m *WsManager) GetClient(userID uint) *client {
	return m.clients[userID]
}

func (m *WsManager) WsListing(conn *websocket.Conn, id uint, hub *Hub) {
	defer conn.Close()
	client := &client{
		conn: conn,
		send: make(chan []byte, 256),
	}
	hub.register <- client
	defer func() { hub.unregister <- client }()
	m.AddClient(id, conn)
	defer m.RemoveClient(id)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		hub.broadcast <- message

		m.SendMessage(id, message)
		log.Printf("received message from client %d: %s", id, message)
	}
}
