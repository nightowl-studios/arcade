package hub

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/gamefactory"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	// hubNameLength for now will just be 4 in length
	hubNameLength int = 4
)

var (
	ErrHubIDNotDefined = errors.New("HubID not found in URL")
)

type HubFactory interface {
	GetHub(r *http.Request, gameFactory gamefactory.GameFactory) *Hub
	SetupRoutes(r *mux.Router)

	CheckIfExists(w http.ResponseWriter, r *http.Request)
	GetNewHubName(w http.ResponseWriter, r *http.Request)
}

// Hub allows for different `GameRouters` to be used within a websocket
// context. As long as these functions are defined, the Hub will
// have access to reading (via HandleMessage) and writing (via `send chan []byte`)
// to the websocket.
type Hub interface {
	// HandleMessage is where the WebsocketHandler should route the messages.
	// Essentially, this function is to become the router for this Game / or API
	HandleMessage(
		messageType int,
		message []byte,
		clientID identifier.Client,
		messageErr error,
	)

	HandleAuthentication(
		w http.ResponseWriter,
		r *http.Request,
		conn *websocket.Conn,
		send chan []byte,
	) (identifier.Client, error)

	// SignalClose is called by the WebsocketClient when the websocket
	// client is about to close.
	SignalClose(caller identifier.Client)

	// Upgrader allows the WebsocketHandler to decide the properties of the
	// websocket upgrade
	Upgrader() *websocket.Upgrader
}

type hubFactory struct {
	registry registry.Registry
}

type hub struct {
	registry    registry.Registry
	gameRouter  game.GameRouter
	gameFactory gamefactory.GameFactory
}

type HubResponse struct {
	Exists bool   `json:"exists"`
	HubID  string `json:"hubID,omitempty"`
}

// GetHubFactory will return a hubFactory
func GetHubFactory(registry registry.Registry) *hubFactory {
	return &hubFactory{
		registry: registry,
	}
}

func (h *hubFactory) GetHub(
	r *http.Request,
	gameFactory gamefactory.GameFactory,
) *hub {
	return &hub{
		registry:    h.registry,
		gameFactory: gameFactory,
	}
}

// SetupRoutes will setup the routes that the hubFactory can respond to
func (h *hubFactory) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/hub", h.GetNewHubName).Methods("GET")
	r.HandleFunc("/hub/{hubID}", h.CheckIfExists).Methods("GET")
}

// GetNewHubName will provide a name of a hub that does not already exist
// in the registry. This function is intended to be used to respond to the
// "Create" button from the front end
func (h *hubFactory) GetNewHubName(w http.ResponseWriter, r *http.Request) {
	// TODO: Need to add some context based cancel timeout later
	var hubName string

	for {
		hubName = identifier.CreateHubName(hubNameLength)
		if !h.registry.CheckIfHubExists(identifier.HubNameStruct{hubName}) {
			// hub does not exist, we can exit
			break
		}
	}

	respondWithJSON(w, http.StatusOK, HubResponse{HubID: hubName})
}

// CheckIfExists will provide a response on whether or not the provided
// HubID exists within the registry. This function is intended to be used
// to respond to the "Join" button from the front end
func (h *hubFactory) CheckIfExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hubID, ok := vars["hubID"]
	if !ok {
		log.Errorf("HubID not found in URL")
		return
	}

	exists := h.registry.CheckIfHubExists(identifier.HubNameStruct{hubID})

	respondWithJSON(w, http.StatusOK, HubResponse{Exists: exists})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// HandleAuthentication is called during the websocket upgrade.
// If there is any authentication handshake, it should be done here.
// The `writePump` will be started prior to calling this function,
// so sending messages via the send channel is the proper way to
// send outgoing messages. As for incoming messages, the `readPump` will
// not be started (to avoid calling the HandleMessage function prior to
// proper authentication) so reading the incoming messages has to be done
// manually
// The WebsocketClient will abort if this function returns a non nil error
func (h *hub) HandleAuthentication(
	w http.ResponseWriter,
	r *http.Request,
	conn *websocket.Conn,
	send chan []byte,
) (client identifier.Client, err error) {
	// no authentication

	vars := mux.Vars(r)
	hubID, ok := vars["hubID"]
	if !ok {
		log.Errorf("%v", ErrHubIDNotDefined)
		return identifier.Client{}, ErrHubIDNotDefined
	}

	// Create an ID
	client = identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{
			UUID: identifier.CreateClientUUID(),
		},
		HubName: identifier.HubNameStruct{
			HubName: hubID,
		},
	}
	h.registry.Register(send, client)
	return
}

func (h *hub) Upgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow all origins to connect
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (h *hub) HandleMessage(
	messageType int,
	message []byte,
	clientID identifier.Client,
	messageErr error,
) {
	// stub
}

func (h *hub) SignalClose(caller identifier.Client) {
	// stub
}
