package hub

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/gamefactory"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	ErrHubIDNotDefined = errors.New("hubID not found in URL")
	ErrNotValidToken   = errors.New("token is not valid")
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

	reg         registry.Registry
	tokenSecret []byte
}

func GetEmptyHub() Hub {
	return &hub{}
}

func GetHub(gameFactory gamefactory.GameFactory) (Hub, error) {
	secret := make([]byte, 60)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	reg := registry.GetRegistryProvider()

	return &hub{
		gameFactory: gameFactory,
		reg:         reg,
		gameRouter:  gameFactory.GetGame("scribble", reg),
		tokenSecret: secret,
	}, nil
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

type JSONWebTokenMessage struct {
	ContainsToken bool   `json:"containsToken"`
	Token         string `json:"token"`
}

type JSONWebToken struct {
	identifier.Client
	jwt.StandardClaims
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

	_, message, err := conn.ReadMessage()
	if err != nil {
		return
	}

	var tokenMessage JSONWebTokenMessage
	err = json.Unmarshal(message, &tokenMessage)
	if err != nil {
		return
	}

	if tokenMessage.ContainsToken {
		client, err := ParseToken(tokenMessage, h.tokenSecret)
		if err == nil {
			// if no error, we can return the client
			// if there is an error, continue and create a new client
			h.RegisterClient(client, send)
			return client, nil
		}
	}

	// no authentication
	client, err = CreateClient(r)

	claims := JSONWebToken{
		Client: client,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(h.tokenSecret)
	tokenMessage = JSONWebTokenMessage{
		Token:         signedString,
		ContainsToken: true,
	}

	message, err = game.MessageBuild("auth", tokenMessage)
	if err != nil {
		log.Errorf("unable to build message: %v", err)
		// continue? I guess they won't be able to re-connect
	}
	send <- message

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

func ParseToken(
	tokenMessage JSONWebTokenMessage,
	tokenSecret []byte,
) (identifier.Client, error) {
	token, err := jwt.ParseWithClaims(
		tokenMessage.Token,
		&JSONWebToken{},
		func(token *jwt.Token) (interface{}, error) {
			return tokenSecret, nil
		},
	)
	if err != nil {
		return identifier.Client{}, err
	}

	if claims, ok := token.Claims.(*JSONWebToken); ok && token.Valid {
		return claims.Client, nil
	}

	return identifier.Client{}, ErrNotValidToken
}

func CreateClient(
	r *http.Request,
) (client identifier.Client, err error) {
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
	return
}
