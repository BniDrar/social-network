package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	scs "socialNetwork/pkg/sessions"
	"socialNetwork/pkg/utils"

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
	GetUserNotification(w http.ResponseWriter, r *http.Request)
	HandleFollowRequestResponse(w http.ResponseWriter, r *http.Request)
	FollowersAndFollowed(w http.ResponseWriter, r *http.Request)
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
		return
	}
	w.Header().Set("Content-Type", "application/json")

	file, fileHeader, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error reading uploaded file"})
		return
	}
	var (
		User        entity.User
		fileContent []byte
	)
	if err != nil {
		User.Avatar = ""
	} else {
		fileContent, User.Avatar, err = utils.ValidateImage(file, fileHeader)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
			return
		}
	}

	err = json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	status, err := u.RegisterService(User)
	if err != nil {
		u.loger.Error.Println(err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	if User.Avatar != "" {
		err = os.WriteFile(User.Avatar, fileContent, 0o444)
		if err != nil {
			u.loger.Error.Println("error while create the avatar", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	w.WriteHeader(status)
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// get user data from request
	var User entity.Credentials
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		u.loger.Error.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	log.Println(User)
	id, err := u.authenticateService(User.Username, User.Password)
	if err != nil {
		if errors.Is(err, config.ErrInvalidCredentials) {
			w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	w.WriteHeader(http.StatusOK)
	u.loger.Info.Println(User, ": is logged in")
}

func (u *user) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Use the RenewToken() method on the current session to change the session
	// ID again. for  session fixation attacks
	err := u.sessionManager.RenewToken(r.Context())
	if err != nil {
		u.loger.Error.Println(err) // that's for registering error in log file
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	u.sessionManager.Remove(r.Context(), "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	u.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	w.WriteHeader(http.StatusOK)
}

func (u *user) Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	target := r.URL.Query().Get("userid")
	id, err := strconv.Atoi(target)
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid user ID"})
		return
	}
	if id == 0 {
		id = r.Context().Value(entity.ContextID).(int)
	}
	status, user, err := u.UserProfile(r.Context(), id)
	if err != nil {
		u.loger.Error.Println(err)
		w.WriteHeader(status)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		u.loger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *user) Follow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	followed := r.URL.Query().Get("followed")
	followedID, err := strconv.Atoi(followed)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid followed ID"})
		return
	}
	status, err := u.FollowService(r.Context(), followedID)
	if err != nil {
		w.WriteHeader(status)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		u.loger.Error.Println(err)
	}
	w.WriteHeader(status)
}

func (u *user) HandleFollowRequestResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var notf entity.Notification
	err := json.NewDecoder(r.Body).Decode(&notf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
	}
	status, err := u.processRequestResponse(r.Context(), notf)
	if err != nil {
		w.WriteHeader(status)
		u.loger.Error.Println(err)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(status)
}

func (u *user) FollowersAndFollowed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	target := r.URL.Query().Get("userid")
	id, err := strconv.Atoi(target)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid user ID"})
		return
	}
	status, follows, err := u.FollowersAndFollowedService(r.Context(), id)
	if err != nil {
		u.loger.Error.Println(err)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "invalid user ID"})
		}
		w.WriteHeader(status)
		return
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(follows)
}

func (u *user) GetUserNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	notifications, status, err := u.userNotificationSerice(r.Context())
	if err != nil {
		u.loger.Error.Println(err)
		w.WriteHeader(status)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(&notifications)
	w.WriteHeader(status)
}
