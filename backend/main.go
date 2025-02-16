package main

import (
	app "socialNetwork/internal/server"
	config "socialNetwork/pkg/config"
	Db "socialNetwork/pkg/db"
	loger "socialNetwork/pkg/loger"
)

func main() {
	// init logger
	Log := loger.NewLogger()
	c, err := config.NewConfig()
	if err != nil {
		Log.Error.Panicln("error while reading config file\n", err)
	}
  // run migrations
	err = Db.RunMigrations()
	if err != nil {
		loger.NewLogger().Error.Fatalln("error while running migrations\n", err)
	}
	// run server
	app.Run(c)
}
