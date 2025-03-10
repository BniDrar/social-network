package user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	scs "socialNetwork/pkg/sessions"

	"socialNetwork/entity"
)

type User interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	Follow(w http.ResponseWriter, r *http.Request)
	Followers(w http.ResponseWriter, r *http.Request)
	Exists(id uint) (bool, error)
}

type user struct {
	serv          *Service
	SessionManger *scs.SessionManager
}

func NewUser(db *sql.DB, s *scs.SessionManager) User {
	repo := NewRepo(db)
	serv := NewService(repo)
	return &user{serv: serv, SessionManger: s}
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	// call service
	token, err, status := u.serv.Login(User)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	// send response and set token in cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	err, status := u.serv.Register(User)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func (u *user) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// app.clientError(w, http.StatusMethodNotAllowed)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Use the RenewToken() method on the current session to change the session
	// ID again. for  session fixation attacks
	// err := u.SessionManger.RenewToken(r.Context())
	// if err != nil {
	// 	http.Error(w, "Method Not Allowed", http.StatusInternalServerError)
	// 	// app.serverError(w, err)
	// 	return
	// }
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	u.SessionManger.Remove(r.Context(), "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	u.SessionManger.Put(r.Context(), "flash", "You've been logged out successfully!")
	// Redirect the user to the application home page.
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (u *user) Profile(w http.ResponseWriter, r *http.Request) {}

func (u *user) Follow(w http.ResponseWriter, r *http.Request) {}

func (u *user) Followers(w http.ResponseWriter, r *http.Request) {}

func (s *user) Exists(id uint) (bool, error) {
	// do something
	return false, nil
}
