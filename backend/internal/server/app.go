package server

import (
	"database/sql"
	"io"
	"log"
	"os"
	"time"

	"socialNetwork/internal/chat"
	"socialNetwork/internal/comment"
	"socialNetwork/internal/group"
	"socialNetwork/internal/user"
	config "socialNetwork/pkg/config"
	"socialNetwork/pkg/database"
	scs "socialNetwork/pkg/sessions"
	"socialNetwork/pkg/store"
	"socialNetwork/pkg/websocket"
)

type App struct {
	sessionManager *scs.SessionManager
	comment.Comment
	chat.Chat
	group.Group
	user.User
	infoLog  *log.Logger
	errorLog *log.Logger
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
	// Use the scs.New() function to initialize a new session manager. Then we
	// configure it to use our MySQL database as the session store, and set a
	// lifetime of 12 hours (so that sessions automatically expire 12 hours
	// after first being created).
	sessionManager := scs.New()
	sessionManager.Store = store.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := NewApp(db)
	app.sessionManager = sessionManager
	server := new(Server)
	// Start listening server
	log.Fatalf("error occured while listening server: %s", server.Run(&cfg.API, app.InitRoutes(cfg)))
}

func NewApp(db *sql.DB) *App {
	hub := websocket.NewHub()
	return &App{
		Comment:  comment.NewComment(db, hub),
		Chat:     chat.NewChat(db, hub),
		Group:    group.NewGroup(db /* we need to add the hub to group*/),
		User:     user.NewUser(db /* we need to add the hub*/),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
