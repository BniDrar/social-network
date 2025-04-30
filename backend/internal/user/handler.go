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
	GetUserPosts(w http.ResponseWriter, r *http.Request)
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

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error parsing form"})
		return
	}

	var user entity.User

	// Required fields
	user.First = r.FormValue("first")
	user.Last = r.FormValue("last")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	user.DateOfBirth = r.FormValue("date_of_birth")
	statusStr := r.FormValue("status")
	
	if statusStr != "" {
		statusUint, err := strconv.ParseUint(statusStr, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "invalid status value"})
			return
		}
		user.Status = uint(statusUint)
	}
	
	// Optional fields
	user.Nickname.String = r.FormValue("nickname")
	user.AboutMe = r.FormValue("about_me")
	file, fileHeader, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error reading uploaded file"})
		return
	}

	var fileContent []byte
	if err == nil {
		fileContent, user.Avatar.NullString.String, err = utils.ValidateImage(file, fileHeader)
		user.Avatar.SetValid(true)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
			return
		}
	}

	status, err := u.RegisterService(user)
	if err != nil {
		u.loger.Error.Println(err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}

	if user.Avatar.Valid {
		err = os.WriteFile(user.Avatar.NullString.String, fileContent, 0o644)
		if err != nil {
			u.loger.Error.Println("error while saving avatar", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error saving avatar"})
			return
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

func (u *user) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil || userId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid query"})
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid query"})
		return
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid query"})
		return
	}

	posts, status, err := u.GetUserPostsService(r.Context(), userId, limit, offset)
	w.WriteHeader(status)
	if err != nil {
		u.loger.Error.Println(err)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		}
		return
	}
	json.NewEncoder(w).Encode(&posts)
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
