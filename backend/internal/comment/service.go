package comment

import "socialNetwork/pkg/websocket"

type Service struct {
	repo *Repository
	hub *websocket.Hub
}



func NewService(repo *Repository, hub *websocket.Hub) *Service {
	return &Service{repo: repo, hub: hub}
}