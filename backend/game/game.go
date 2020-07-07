// package game defines the type of handlers a game needs to have
package game

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

// Message is the 'envelope' that all game messages should contain
// This is the top level message that the GameHandler will receive
// and it is necessary for the GameHandler as it needs the information
// for routing
type Message struct {
	API     string          `json:"api"`
	Payload json.RawMessage `json:"payload"`
}

// MessageBuild will take the payload and api, and build the
// message that can be sent back through to the websockets
func MessageBuild(api string, payload interface{}) (json.RawMessage, error) {

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("unable to marshal json: %v", err)
		return nil, err
	}

	retJson, err := json.Marshal(Message{
		API:     api,
		Payload: bytePayload,
	})
	if err != nil {
		log.Errorf("unable to marshal json: %v", err)
		return nil, err
	}

	return retJson, nil
}

type GameRouter interface {
	RouteMessage(
		messageType int,
		message []byte,
		clientID identifier.Client,
		messageErr error,
		reg registry.Registry,
	)

	NewClient(
		clientID identifier.Client,
		reg registry.Registry,
	)

	ClientQuit(
		clientID identifier.Client,
		reg registry.Registry,
	)

	RouterName() string
}

type GameHandler interface {

	// HandleInteraction will be given the tools it needs to handle
	// any interaction
	HandleInteraction(
		message json.RawMessage,
		caller identifier.Client,
		registry registry.Registry,
	)

	NewClient(
		clientID identifier.Client,
		reg registry.Registry,
	)

	ClientQuit(
		clientID identifier.Client,
		reg registry.Registry,
	)

	// Name needs to return a unique name of this GameHandler
	// This return will be used for routing
	Name() string
}

func CreateGameHandlersMap(handlers ...GameHandler) map[string]GameHandler {
	handlerMap := make(map[string]GameHandler)
	for _, handler := range handlers {
		handlerMap[handler.Name()] = handler
	}
	return handlerMap
}
