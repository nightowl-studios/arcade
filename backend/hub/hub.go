package hub

import (
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

var (
	ErrHubIDNotDefined = errors.New("HubID not found in URL")
)

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

	// Upgrader allows the WebsocketHandler to decide the properties of the
	// websocket upgrade
	Upgrader() *websocket.Upgrader

	RegisterClient(clientID identifier.Client, send chan []byte)
	UnregisterClient(clientID identifier.Client) (hubEmpty bool)
}

type hub struct {
	// we need the gameFactory since the user can choose after creating
	// a hub, which game they'd like to play
	gameFactory gamefactory.GameFactory
	gameRouter  game.GameRouter

	reg registry.Registry
}

// hubClient should only store basic information about the client
// and the send channel
type hubClient struct {
	send     (chan []byte)
	nickname string
}

func GetEmptyHub() Hub {
	return &hub{}
}

func GetHub(gameFactory gamefactory.GameFactory) Hub {
	return &hub{
		gameFactory: gameFactory,
		reg:         registry.GetRegistryProvider(),
		gameRouter:  gameFactory.GetGame("scribble"),
	}
}

func (h *hub) RegisterClient(clientID identifier.Client, send chan []byte) {
	h.reg.Register(send, clientID)
	h.gameRouter.NewClient(clientID, h.reg)
}

func (h *hub) UnregisterClient(
	clientID identifier.Client,
) (hubEmpty bool) {
	hubEmpty = h.reg.Unregister(clientID)
	if hubEmpty != true {
		h.gameRouter.ClientQuit(clientID, h.reg)
	}
	return
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

	h.RegisterClient(client, send)
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
	h.gameRouter.RouteMessage(
		messageType,
		message,
		clientID,
		messageErr,
		h.reg,
	)
}
