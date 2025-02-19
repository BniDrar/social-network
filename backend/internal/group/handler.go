package group

import (
	"database/sql"
	"net/http"
)

type Group interface {
	Group(w http.ResponseWriter, r *http.Request)
}

type group struct {
	Service *Service
}

func NewGroup(db *sql.DB) Group {
	repo := NewRepo(db)
	serv := NewService(repo)
	return &group{Service: serv}
}

func (g *group) Group(w http.ResponseWriter, r *http.Request) {}