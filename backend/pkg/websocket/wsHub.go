package websocket

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type MessageType int

const (
	PrivateMessage MessageType = iota
	GroupMessage
	BroadcastMessage
	NotificationMessage
)

type Message struct {
	Type         MessageType
	UserID       uint
	GroupID      uint
	GroupMembers []uint
	Data         []byte
}

type Client struct {
	userID uint
	conn   *websocket.Conn
	send   chan []byte
}

type Hub struct {
	// WebSocket connections
	Clients map[uint]*Client
	// Channels for different types of messages
	private      chan Message
	group        chan Message
	broadcast    chan Message
	notification chan Message
	// Channel for registering/unregistering Clients
	register   chan *Client
	unregister chan *Client
	// Mutex for thread safety
	mu sync.Mutex
	// Configuration
	bufferSize int
}

func NewHub(bufferSize int) *Hub {
	if bufferSize <= 0 {
		bufferSize = 1024 // Default buffer size
	}
	return &Hub{
		Clients:      make(map[uint]*Client),
		private:      make(chan Message),
		group:        make(chan Message),
		broadcast:    make(chan Message),
		notification: make(chan Message),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		bufferSize:   bufferSize,
	}
}

func (h *Hub) AddClient(userID uint, conn *websocket.Conn) {
	client := &Client{
		userID: userID,
		conn:   conn,
		send:   make(chan []byte, h.bufferSize),
	}
	h.register <- client

	// Start ping/pong handling
	go h.handlePingPong(client)
	// Start message writing
	go h.writeMessage(client)
}

func (h *Hub) handlePingPong(client *Client) {
	// Set ping handler
	client.conn.SetPingHandler(func(appData string) error {
		return client.conn.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(time.Second))
	})

	// Set pong handler
	lastPongTime := time.Now() // Initialize with current time
	client.conn.SetPongHandler(func(appData string) error {
		lastPongTime = time.Now()
		return nil
	})

	// Send ping every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if time.Since(lastPongTime) > 20*time.Second {
				log.Printf("No pong from client %d, closing connection", client.userID)
				h.unregister <- client
				return
			}
			if err := client.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second)); err != nil {
				log.Printf("Error sending ping to client %d: %v", client.userID, err)
				h.unregister <- client
				return
			}
		}
	}
}

func (h *Hub) writeMessage(client *Client) {
	for message := range client.send {
		err := client.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing message to client %d: %v", client.userID, err)
			h.unregister <- client
			return
		}
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.Clients[client.userID] = client
			h.mu.Unlock()
			log.Printf("Client %d registered", client.userID)

		case client := <-h.unregister:
			log.Println("here")
			h.mu.Lock()
			if _, ok := h.Clients[client.userID]; ok {
				delete(h.Clients, client.userID)
				close(client.send)
				client.conn.Close() // Close the connection
				log.Printf("Client %d unregistered", client.userID)
			}
			h.mu.Unlock()

		case message := <-h.private:
			h.mu.Lock()
			if client, ok := h.Clients[message.UserID]; ok {
				jsonData, err := json.Marshal(message)
				if err != nil {
					log.Printf("Error marshaling message: %v", err)
					continue
				}
				select {
				case client.send <- jsonData:
				default:
					log.Printf("Client %d send buffer is full", message.UserID)
				}
			}
			h.mu.Unlock()

		case message := <-h.group:
			jsonData, err := json.Marshal(message)
			if err != nil {
				log.Printf("Error marshaling message: %v", err)
				continue
			}
			h.mu.Lock()
			for _, userID := range message.GroupMembers {
				if client, exists := h.Clients[userID]; exists {
					select {
					case client.send <- jsonData:
					default:
						log.Printf("Client %d send buffer is full", userID)
					}
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.Lock()
			jsonData, err := json.Marshal(message)
			if err != nil {
				log.Printf("Error marshaling message: %v", err)
				continue
			}
			for userID, client := range h.Clients {
				if userID != message.UserID { // Don't send to sender

					select {
					case client.send <- jsonData:
					default:
						log.Printf("Client %d send buffer is full", userID)
					}
				}
			}
			h.mu.Unlock()

		case message := <-h.notification:
			h.mu.Lock()
			if client, ok := h.Clients[message.UserID]; ok {
				jsonData, err := json.Marshal(message)
				if err != nil {
					log.Printf("Error marshaling message: %v", err)
					continue
				}
				select {
				case client.send <- jsonData:
				default:
					log.Printf("Client %d send buffer is full", message.UserID)
				}
			}
			h.mu.Unlock()
		}
	}
}

// SendPrivateMessage sends a message to a specific user
func (h *Hub) SendPrivateMessage(userID uint, message []byte) {
	h.private <- Message{
		Type:   PrivateMessage,
		UserID: userID,
		Data:   message,
	}
}

// SendGroupMessage sends a message to all members of a group
func (h *Hub) SendGroupMessage(groupID uint, groupMembers []uint, message []byte) {
	h.group <- Message{
		Type:         GroupMessage,
		GroupID:      groupID,
		GroupMembers: groupMembers,
		Data:         message,
	}
}

// BroadcastMessage sends a message to all connected Clients except the sender
func (h *Hub) BroadcastMessage(senderID uint, message []byte) {
	h.broadcast <- Message{
		Type:   BroadcastMessage,
		UserID: senderID,
		Data:   message,
	}
}

// SendNotification sends a notification to a specific user
func (h *Hub) SendNotification(userID uint, message []byte) {
	h.notification <- Message{
		Type:   NotificationMessage,
		UserID: userID,
		Data:   message,
	}
}

func (h *Hub) GetOnlineUsers() []uint {
	onlineUsers := make([]uint, 0)
	for userID := range h.Clients {
		onlineUsers = append(onlineUsers, userID)
	}
	return onlineUsers
}

func (h *Hub) IsOnline(userId uint) bool {
	return h.Clients[userId] != nil
}

func (h *Hub) Unregister(userID uint) {
	h.unregister <- h.Clients[userID]
}
