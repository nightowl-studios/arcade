// package scribble will define a `WebsocketHandler` to handle the backend for
// the scribble game
package scribble

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/generic/chat"
	"github.com/bseto/arcade/backend/game/hubapi"
	"github.com/bseto/arcade/backend/game/scribble/handler/addition"
	"github.com/bseto/arcade/backend/game/scribble/handler/draw"
	"github.com/bseto/arcade/backend/game/scribble/handler/echo"
	"github.com/bseto/arcade/backend/game/scribble/handler/gamemaster"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name string = "scribble"
)

type State uint32

// States
const (
	Lobby State = iota
	PlayerSelectTopic
	PlayTime
	ScoreTime
	Results
)

type Router struct {
	// a simple ad-hoc pub/sub structure
	handlers map[string][]game.GameHandler
}

func GetScribbleRouter(reg registry.Registry) game.GameRouter {
	var handlers map[string][]game.GameHandler
	if reg != nil {
		handlers = game.CreateGameHandlersMap(
			echo.Get(),
			addition.Get(),
			hubapi.Get(),
			chat.Get(),
			draw.Get(),
			gamemaster.Get(reg),
		)
	}

	return &Router{
		handlers: handlers,
	}
}

func (r *Router) RouterName() string {
	return name
}

// NOTE, change this signature to match game.go::GameRouter
// HandleMessage is the router to GameHandlers
func (r *Router) RouteMessage(
	messageType int,
	message []byte,
	clientID identifier.Client,
	messageErr error,
	reg registry.Registry,
) {
	var msg game.Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal the message: %v", err)
	}

	handlers, ok := r.handlers[msg.API]
	if !ok {
		log.Errorf("unable to find handler for: %v", msg.API)
		return
	}

	// cheap pub/sub using map of []handlers
	for _, handler := range handlers {
		handler.HandleInteraction(
			msg.Payload,
			clientID,
			reg,
		)
	}

	return
}

// NewClient will just tell any handler - if they care - that there is a new client
func (r *Router) NewClient(
	clientID identifier.Client,
	reg registry.Registry,
) {
	for _, handlers := range r.handlers {
		for _, handler := range handlers {
			handler.NewClient(clientID, reg)
		}
	}
}

// ClientQuit will just tell any handler - if they care - that a client quit
func (r *Router) ClientQuit(
	clientID identifier.Client,
	reg registry.Registry,
) {
	for _, handlers := range r.handlers {
		for _, handler := range handlers {
			handler.ClientQuit(clientID, reg)
		}
	}
}
