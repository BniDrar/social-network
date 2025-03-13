package group

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type Group interface {
	GetGroups(w http.ResponseWriter, r *http.Request)
}

type group struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

func NewGroup(dep *config.Dependencies) Group {
	return &group{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (g *group) GetGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Method not allowed"})
		return
	}
	// get the limit and offset from the request body
	var requestBody struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	limit := requestBody.Limit
	offset := requestBody.Offset
	if limit <= 0 || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid limit or offset"})
		return
	}

	groups, err := g.GetGroupsByUser(r.Context(), limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}
