package server

import (
	"net/http"

	"socialNetwork/pkg/config"
)

type Route struct {
	Path    string
	handler http.HandlerFunc
	Role    uint
}

func (app *App) InitRoutes(conf *config.Conf) *http.ServeMux {
	mux := http.NewServeMux()
	routes := app.createRoutes()
	for _, route := range routes {
		mux.Handle(route.Path, route.handler)
	}
	return mux
}

func (app *App) createRoutes() []Route {
	return []Route{
		{
			Path:    "/api/register",
			handler: app.Register,
			Role:    0,
		},
	}
}
