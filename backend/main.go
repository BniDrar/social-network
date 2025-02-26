package main

import (
	app "socialNetwork/internal/server"
	config "socialNetwork/pkg/config"
	loger "socialNetwork/pkg/loger"
)

func main() {
	// init logger
	Log := loger.NewLogger()
	c, err := config.NewConfig()
	if err != nil {
		Log.Error.Panicln("error while reading config file\n", err)
	}
	// run server
	app.Run(c)
}
