package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	scs "socialNetwork/pkg/sessions"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type user struct {
	db             *sql.DB
	hub            *websocket.Hub
	loger          *loger.CstmLogger
	sessionManager *scs.SessionManager
}

type User interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	Follow(w http.ResponseWriter, r *http.Request)
	Followers(w http.ResponseWriter, r *http.Request)
	DeleteUserByNickName(Nickname string) error
	IsUserExist(id uint) (bool, error)
}

func NewUser(dep *config.Dependencies) User {
	return &user{
		db:             dep.DB,
		hub:            dep.Hub,
		loger:          dep.Loger,
		sessionManager: dep.SessionManager,
	}
}

func (u *user) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Method not allowed"})
		return
	}
	User := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		u.loger.Info.Println("error here ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	status, err := u.RegisterService(User)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get user data from request
	var User entity.Credentials
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		u.loger.Error.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	id, err := u.authenticateService(User.Username, User.Password)
	if err != nil {
		if errors.Is(err, config.ErrInvalidCredentials) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		} else {
			u.loger.Error.Println(err) // that's for registering error in log file
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	err = u.sessionManager.RenewToken(r.Context())
	if err != nil {
		u.loger.Error.Println(err) // that's for registering error in log file
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	u.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *user) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		u.loger.Error.Println("Method Not allowed")
		// app.clientError(w, http.StatusMethodNotAllowed)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Use the RenewToken() method on the current session to change the session
	// ID again. for  session fixation attacks
	err := u.sessionManager.RenewToken(r.Context())
	if err != nil {
		http.Error(w, "Method Not Allowed", http.StatusInternalServerError)
		// app.serverError(w, err)
		return
	}
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	u.sessionManager.Remove(r.Context(), "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	u.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")
	// Redirect the user to the application home page.
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}

func (u *user) Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		u.loger.Error.Println("the method used is not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	nickname := r.URL.Query().Get("nickname")
	status, user, err := u.UserProfile(r.Context(), nickname)
	if err != nil {
		u.loger.Error.Println(err)
		http.Error(w, err.Error(), status)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		u.loger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *user) Follow(w http.ResponseWriter, r *http.Request) {}

func (u *user) Followers(w http.ResponseWriter, r *http.Request) {}

