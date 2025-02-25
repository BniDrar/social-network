package server

import (
	"database/sql"
	"io"
	"log"
	"os"
	"socialNetwork/internal/chat"
	"socialNetwork/internal/comment"
	"socialNetwork/internal/group"
	"socialNetwork/internal/user"
	config "socialNetwork/pkg/config"
	"socialNetwork/pkg/database"
	"socialNetwork/pkg/websocket"
)

type App struct {
	comment.Comment
	chat.Chat
	group.Group
	user.User
}

func Run(cfg *config.Conf) {
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Printf("cannot create log file: %v", err)
	}
	defer file.Close()
	logWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logWriter)

	// Prepare database
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("error occured while connecting database: %s", err.Error())
		return
	}
	// Close connection database
	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal("can't close connection db, err:", err)
		} else {
			log.Println("db closed")
		}
	}()
	app := NewApp(db)
	server := new(Server)
	// Start listening server
	log.Fatalf("error occured while listening server: %s", server.Run(&cfg.API, app.InitRoutes(cfg)))
}

func NewApp(db *sql.DB) *App {
	hub := websocket.NewHub()
	return &App{
		Comment: comment.NewComment(db, hub),
		Chat:    chat.NewChat(db, hub),
		Group:   group.NewGroup(db /* we need to add the hub to group*/),
		User:    user.NewUser(db /* we need to add the hub*/),
	}
}
