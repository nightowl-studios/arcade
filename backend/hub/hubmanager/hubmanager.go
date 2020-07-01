package hubmanager

import (
	"encoding/json"
	"net/http"

	"github.com/bseto/arcade/backend/game/gamefactory"
	"github.com/bseto/arcade/backend/hub"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/gorilla/mux"
)

const (
	// hubNameLength for now will just be 4 in length
	hubNameLength int = 4
)

type HubManager interface {
	GetHub(r *http.Request, gameFactory gamefactory.GameFactory) (hub.Hub, error)
	SetupRoutes(r *mux.Router)

	CheckIfExists(w http.ResponseWriter, r *http.Request)
	GetNewHubName(w http.ResponseWriter, r *http.Request)
}

type hubManager struct {
	hubs map[identifier.HubNameStruct]hub.Hub
}

type HubResponse struct {
	Exists bool   `json:"exists"`
	HubID  string `json:"hubID,omitempty"`
}

// GetHubFactory will return a hubManager
func GetHubManager() *hubManager {
	return &hubManager{
		hubs: make(map[identifier.HubNameStruct]hub.Hub),
	}
}

func (h *hubManager) WebsocketClose(clientID identifier.Client) {
	hubInstance := h.hubs[clientID.HubName]
	hubEmpty := hubInstance.UnregisterClient(clientID)
	if hubEmpty {
		delete(h.hubs, clientID.HubName)
	}
}

func (h *hubManager) GetHub(
	r *http.Request,
	gameFactory gamefactory.GameFactory,
) (hub.Hub, error) {

	vars := mux.Vars(r)
	hubID, ok := vars["hubID"]
	if !ok {
		log.Errorf("%v", hub.ErrHubIDNotDefined)
		return hub.GetEmptyHub(), hub.ErrHubIDNotDefined
	}
	hubStruct := identifier.HubNameStruct{
		HubName: hubID,
	}

	hubInstance, ok := h.hubs[hubStruct]
	if !ok {
		hubInstance = hub.GetHub(gameFactory)
		h.hubs[hubStruct] = hubInstance
	}

	return hubInstance, nil
}

// SetupRoutes will setup the routes that the hubManager can respond to
func (h *hubManager) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/hub", h.GetNewHubName).Methods("GET")
	r.HandleFunc("/hub/{hubID}", h.CheckIfExists).Methods("GET")
}

// GetNewHubName will provide a name of a hub that does not already exist
// in the registry. This function is intended to be used to respond to the
// "Create" button from the front end
func (h *hubManager) GetNewHubName(w http.ResponseWriter, r *http.Request) {
	// TODO: Need to add some context based cancel timeout later
	var hubName string

	for {
		hubName = identifier.CreateHubName(hubNameLength)
		if _, ok := h.hubs[identifier.HubNameStruct{hubName}]; !ok {
			// hub does not exist, we have found a valid name
			break
		}
	}

	respondWithJSON(w, http.StatusOK, HubResponse{HubID: hubName})
}

// CheckIfExists will provide a response on whether or not the provided
// HubID exists within the registry. This function is intended to be used
// to respond to the "Join" button from the front end
func (h *hubManager) CheckIfExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hubID, ok := vars["hubID"]
	if !ok {
		log.Errorf("HubID not found in URL")
		return
	}

	_, exists := h.hubs[identifier.HubNameStruct{hubID}]
	respondWithJSON(w, http.StatusOK, HubResponse{Exists: exists})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
