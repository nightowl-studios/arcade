package hub

import (
	"encoding/json"
	"net/http"

	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/gorilla/mux"
)

const (
	// hubNameLength for now will just be 4 in length
	hubNameLength int = 4
)

type Hub struct {
	registry registry.Registry
}

type HubResponse struct {
	Exists bool   `json:"exists"`
	HubID  string `json:"hubID,omitempty"`
}

func GetHub(registry registry.Registry) *Hub {
	return &Hub{
		registry: registry,
	}
}

func (l *Hub) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/hub", l.GetNewHubName).Methods("GET")
	r.HandleFunc("/hub/{hubID}", l.CheckIfExists).Methods("GET")
}

// GetNewHubName will provide a name of a hub that does not already exist
// in the registry. This function is intended to be used to respond to the
// "Create" button from the front end
func (l *Hub) GetNewHubName(w http.ResponseWriter, r *http.Request) {
	// TODO: Need to add some context based cancel timeout later
	var hubName string

	for {
		hubName = identifier.CreateHubName(hubNameLength)
		if !l.registry.CheckIfHubExists(identifier.HubNameStruct{hubName}) {
			// hub does not exist, we can exit
			break
		}
	}

	respondWithJSON(w, http.StatusOK, HubResponse{HubID: hubName})
}

// CheckIfExists will provide a response on whether or not the provided
// HubID exists within the registry. This function is intended to be used
// to respond to the "Join" button from the front end
func (l *Hub) CheckIfExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hubID, ok := vars["hubID"]
	if !ok {
		log.Errorf("HubID not found in URL")
		return
	}

	exists := l.registry.CheckIfHubExists(identifier.HubNameStruct{hubID})

	respondWithJSON(w, http.StatusOK, HubResponse{Exists: exists})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
