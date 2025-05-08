package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"socialNetwork/internal/server"
	config "socialNetwork/pkg/config"
	"socialNetwork/pkg/database"
	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
	"socialNetwork/pkg/store"
	"socialNetwork/pkg/websocket"
)

func main() {
	// init logger
	loger := loger.NewLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		loger.Error.Panicln("error while reading config file\n", err)
	}

	db, err := database.InitDB(cfg.Database)
	if err != nil {
		loger.Error.Panicf("error occured while connecting database: %s", err.Error())
	}
	defer db.Close()

	// Use the scs.New() function to initialize a new session manager. Then we
	// configure it to use our MySQL database as the session store, and set a
	// lifetime of 12 hours (so that sessions automatically expire 12 hours
	// after first being created).
	sessionManager := scs.New()
	sessionManager.Store = store.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	hub := websocket.NewHub(1024)
	hub.Run()
	Dep := &config.Dependencies{
		SessionManager: sessionManager,
		DB:             db,
		Loger:          loger,
		/*Legislation is the process or result of enrolling, enacting, or promulgating laws by a legislature, parliament, or analogous governing body.*/
		//hub := websocket.NewHub() // need console legislation
		Hub: hub,
	}

	app := server.NewApp(Dep)

	server := http.Server{
		Addr:    ":" + cfg.API.Port,
		Handler: app.InitRoutes(cfg),
	}

	// Start listening server
	loger.Info.Printf("\033[32mServer is running...🚀\n\tLink: 🌐 http://%s:%s", cfg.API.Host, cfg.API.Port)
	serverErrChan := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			serverErrChan <- err
		} else {
			serverErrChan <- nil
		}
	}()

	// Create a channel to capture OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Println("\nserver shut down gracefully:", sig)
		// Gracefully shutdown
		if err := server.Close(); err != nil {
			log.Fatal("Server Close failed:", err)
		}
	case err := <-serverErrChan:
		if err != nil {
			log.Fatal("Server crashed:", err)
		} else {
			log.Println("Server shutdown gracefully")
		}
	}
}
