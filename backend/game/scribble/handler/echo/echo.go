package echo

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name string = "echo"
)

type EchoMessage struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message string
}

type Echo struct{}

func Get() *Echo {
	return &Echo{}
}

// HandleInteraction will echo with a flavour :D
func (e *Echo) HandleInteraction(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var msg EchoMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal message: %v", err)
		return
	}

	newMessage, err := json.Marshal(msg.Message + " Gordon")
	if err != nil {
		log.Errorf("unable to marshal message: %v", err)
		return
	}

	response := game.Message{
		API:     name,
		Payload: newMessage,
	}

	responseMessage, err := json.Marshal(response)
	if err != nil {
		log.Errorf("unable to marshal the full response: %v", err)
		return
	}

	registry.SendToCaller(caller, responseMessage)
}

func (e *Echo) Name() string {
	return name
}
