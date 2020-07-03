// package scribble will define a `WebsocketHandler` to handle the backend for
// the scribble game
package scribble

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/scribble/handler/addition"
	"github.com/bseto/arcade/backend/game/scribble/handler/echo"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name string = "scribble"
)

type Router struct {
	handlers map[string]game.GameHandler
}

func GetScribbleRouter() *Router {
	handlers := game.CreateGameHandlersMap(
		echo.Get(),
		addition.Get(),
	)

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

	handler, ok := r.handlers[msg.API]
	if !ok {
		log.Errorf("unable to find handler for: %v", msg.API)
		return
	}

	handler.HandleInteraction(
		msg.Payload,
		clientID,
		reg,
	)

	return
}
