// package scribble will define a `WebsocketHandler` to handle the backend for
// the scribble game
package scribble

import (
	"encoding/json"
	"net/http"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/scribble/handler/addition"
	"github.com/bseto/arcade/backend/game/scribble/handler/echo"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/gorilla/websocket"
)

type API struct {
	handlers map[string]game.GameHandler
	registry registry.Registry
}

func GetScribbleAPI(reg registry.Registry) *API {

	handlers := game.CreateGameHandlers(
		echo.Get(),
		addition.Get(),
	)

	return &API{
		handlers: handlers,
		registry: reg,
	}
}

// HandleMessage is the router to GameHandlers
func (a *API) HandleMessage(
	messageType int,
	message []byte,
	clientID identifier.Client,
	messageErr error,
) {
	var msg game.Message

	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal the message: %v", err)
	}

	handler, ok := a.handlers[msg.API]
	if !ok {
		log.Errorf("unable to find handler for: %v", msg.API)
		return
	}

	handler.HandleInteraction(
		msg.Payload,
		clientID,
		a.registry,
	)

	return
}

func (a *API) HandleAuthentication(
	w http.ResponseWriter,
	r *http.Request,
	conn *websocket.Conn,
	send chan []byte,
) (client identifier.Client, err error) {
	// no authentication

	// Create an ID
	client = identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{
			UUID: identifier.CreateClientUUID(),
		},
		HubName: identifier.HubNameStruct{
			HubName: r.URL.Path,
		},
	}
	a.registry.Register(send, client)
	return
}

func (a *API) SignalClose(caller identifier.Client) {
	a.registry.Unregister(caller)
	return
}

func (a *API) Upgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow all origins to connect
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
