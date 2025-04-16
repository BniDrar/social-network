package group

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type Group interface {
	GetUserGroups(w http.ResponseWriter, r *http.Request)
	GetGroupById(w http.ResponseWriter, r *http.Request)
	CreateGroup(w http.ResponseWriter, r *http.Request)
	GetAllGroups(w http.ResponseWriter, r *http.Request)
	GetGroupMembers(w http.ResponseWriter, r *http.Request)
	RequestToJoinGroup(w http.ResponseWriter, r *http.Request)
	AcceptRequestToJoinGroup(w http.ResponseWriter, r *http.Request)
}

type group struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

func NewGroup(dep *config.Dependencies) Group {
	return &group{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

// this handler is used to get all groups
func (g *group) GetUserGroups(w http.ResponseWriter, r *http.Request) {
	// get the limit and offset from the request body
	var requestBody struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
	log.Println("request body", r.Body)
	log.Println("request body decoded", requestBody)
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	limit := requestBody.Limit
	offset := requestBody.Offset
	log.Println("Limit:", limit)
	log.Println("Offset:", offset)
	if limit <= 0 || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid limit or offset"})
		return
	}

	groups, err := g.GetGroupsByUserService(r.Context(), limit, offset)
	if err != nil {
		log.Println("group err 0", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

// this handler is used to get the group by id
func (g *group) GetGroupById(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || groupID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid group ID"})
		return
	}
	group, err := g.GetGroupByIdService(r.Context(), groupID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}

// this handler is used to create a group
func (g *group) CreateGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// get the group data from the request body
	var group entity.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	// create the group in the database
	createdGroup, err := g.CreateGroupService(r.Context(), group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	// send the created group to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdGroup)
}

// this handler is used to get all groups
func (g *group) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	g.loger.Info.Println("In Get All Groups")
	// get the limit and offset from the request body
	var requestBody struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		// Type   int `json:"type"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	limit := requestBody.Limit
	offset := requestBody.Offset
	// typeGroup := requestBody.Type
	if limit <= 0 || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid limit or offset"})
		return
	}
	groups, err := g.GetAllGroupsService(r.Context(), limit, offset, entity.RealGroup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

// this handler is used to get the group members
func (g *group) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || groupID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid group ID"})
		return
	}
	group, err := g.GetGroupMembersService(r.Context(), groupID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}

// this handler is used to request to join a group
func (g *group) RequestToJoinGroup(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || groupID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid group ID"})
		return
	}
	// get the user ID from the request body
	var requestBody struct {
		UserID int `json:"user_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	// request to join the group
	err = g.RequestToJoinGroupService(r.Context(), groupID, requestBody.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}

// this handler is used to accept a request to join a group
func (g *group) AcceptRequestToJoinGroup(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || groupID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid group ID"})
		return
	}
	// get the user ID from the request body
	var requestBody struct {
		UserID int `json:"user_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	// accept the request to join the group
	err = g.AcceptRequestToJoinGroupService(r.Context(), groupID, requestBody.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}
