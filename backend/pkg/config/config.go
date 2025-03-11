package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
	"socialNetwork/pkg/websocket"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	// Add a new ErrInvalidCredentials error. We'll use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type (
	Conf struct {
		API            API      `json:"api"`
		Database       Database `json:"database"`
		TestDatabase   Database `json:"testdatabase"`
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

	TestDB struct {
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
