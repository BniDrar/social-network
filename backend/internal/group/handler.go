package group

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type Group interface {
	GetGroups(w http.ResponseWriter, r *http.Request)
	GetGroupById(w http.ResponseWriter, r *http.Request)
	CreateGroup(w http.ResponseWriter, r *http.Request)
	GetAllGroups(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	VoteEvent(w http.ResponseWriter, r *http.Request)
	GetEvent(w http.ResponseWriter, r *http.Request)
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
func (g *group) GetGroups(w http.ResponseWriter, r *http.Request) {
	// get the limit and offset from the request body
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		g.loger.Error.Println("Invalid limit:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		g.loger.Error.Println("Invalid offset:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid offset"})
		return
	}

	g.loger.Info.Println("Limit:", limit, "Offset:", offset)
	if limit <= 0 || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid limit or offset"})
		return
	}

	groups, err := g.GetGroupsByUserService(r.Context(), limit, offset)
	if err != nil {
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

func (g *group) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// get the event data from the request body
	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	// create the event in the database
	eventId, status, err := g.CreateEventService(r.Context(), event)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	// send the created event to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct{
		EventId int `json:"event_id"`
	}{
		EventId: eventId,
	})
}

func (g *group) VoteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// get the event data from the request body
	var vote entity.Engagement
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid request body"})
		return
	}
	// create the event in the database
	eventId, status, err := g.VoteEventService(r.Context(), vote)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	// send the created event to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		EventID int `json:"event_id"`
	}{
		EventID: eventId,
	})
}


func (g *group) GetEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// get the event data from the request body
	eventID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || eventID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Invalid event ID"})
		return
	}
	// create the event in the database
	event, status, err := g.GetEventService(r.Context(), eventID)
	if err != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: err.Error()})
		return
	}
	// send the created event to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(event)
}