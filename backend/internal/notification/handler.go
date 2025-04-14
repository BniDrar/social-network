package notification

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
)

type Notification struct {
	db    *sql.DB
	loger *loger.CstmLogger
}

func NewNotification(dep *config.Dependencies) *Notification {
	return &Notification{
		db:    dep.DB,
		loger: dep.Loger,
	}
}

func (n *Notification) ReactToNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var notf entity.Notification
	err := json.NewDecoder(r.Body).Decode(&notf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
	}
}
