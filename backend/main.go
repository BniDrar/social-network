package main

import (
	"log"
	app "socialNetwork/internal/server"
	config "socialNetwork/pkg/config"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Panic("error while getting configurations\n", err)
	}
	app.Run(c)
}
