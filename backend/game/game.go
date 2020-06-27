// package game defines the type of handlers a game needs to have
package game

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

type GameHandler interface {

	// HandleInteraction will be given the tools it needs to handle
	// any interaction
	HandleInteraction(
		message json.RawMessage,
		caller identifier.Client,
		registry registry.Registry,
	)

	// Name needs to return a unique name of this GameHandler
	// This return will be used for routing
	Name() string
}

func CreateGameHandlers(handlers ...GameHandler) map[string]GameHandler {
	handlerMap := make(map[string]GameHandler)
	for _, handler := range handlers {
		handlerMap[handler.Name()] = handler
	}
	return handlerMap
}
