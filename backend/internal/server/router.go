package server

import (
	"fmt"
	"net/http"

	"socialNetwork/pkg/config"
)

type Route struct {
	Path    string
	handler http.HandlerFunc
	Role    uint
}

const (
	Auth uint = iota
	User
)

func (app *App) InitRoutes(conf *config.Conf) http.Handler {
	mux := http.NewServeMux()
	routes := app.createRoutes()
	for _, route := range routes {
		if requireLogin(route.Role) {
			fmt.Println("require login")
			mux.Handle(route.Path, app.sessionManager.LoadAndSave(app.authenticate(app.requireAuthentication(http.HandlerFunc(route.handler)))))
		} else {
			fmt.Println("doesnt' require logging")
			mux.Handle(route.Path, route.handler)
		}
	}
	return mux
}

func (app *App) createRoutes() []Route {
	return []Route{
		{
			Path:    "/api/register",
			handler: app.Register,
			Role:    User,
		},
		{
			Path:    "/api/login",
			handler: app.Login,
			Role:    Auth,
		},
		{
			Path:    "/api/logout",
			handler: app.Logout,
			Role:    User,
		},
	}
}

func requireLogin(Auth uint) bool {
	return Auth == User
}
