package server

import (
	"socialNetwork/internal/chat"
	"socialNetwork/internal/comment"
	"socialNetwork/internal/group"
	"socialNetwork/internal/user"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
)

type App struct {
	SessionManager *scs.SessionManager
	comment.Comment
	chat.Chat
	group.Group
	user.User
	Loger *loger.CstmLogger
}

func NewApp(dep *config.Dependencies) *App {
	return &App{
		SessionManager: dep.SessionManager,
		Loger:          dep.Loger,
		Comment:        comment.NewComment(dep),
		Chat:           chat.NewChat(dep),
		Group:          group.NewGroup(dep /* we need to add the hub to group*/),
		User:           user.NewUser(dep /* we need to add the hub*/),
	}
}
