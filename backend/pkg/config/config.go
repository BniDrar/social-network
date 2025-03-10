package config

import (
	"database/sql"
	"encoding/json"
	"os"

	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
	"socialNetwork/pkg/websocket"
)

type (
	Conf struct {
		API            API      `json:"api"`
		Database       Database `json:"database"`
		SessionManager *scs.SessionManager
	}

	API struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Database struct {
		Driver    string `json:"driver"`
		FileName  string `json:"fileName"`
		SchemeDir string `json:"schemeDir"`
	}

	Dependencies struct {
		Loger          *loger.CstmLogger
		DB             *sql.DB
		SessionManager *scs.SessionManager
		Hub            *websocket.Hub
	}
)

func NewConfig() (*Conf, error) {
	var newConfig Conf
	file, err := os.Open("pkg/config/config.json")
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(file).Decode(&newConfig); err != nil {
		return nil, err
	}

	return &newConfig, nil
}
