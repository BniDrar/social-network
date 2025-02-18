package server

import (
	config "socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
)

const secret string = "Forum01Oujda"

func Run(cfg *config.Conf)  {
	loger.NewLogger().Info.Println("Starting server on port: ", cfg.API.Port)
}
