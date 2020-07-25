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

type Handler struct {
	name string
}

func Get() *Handler {
	return &Handler{
		name: name,
	}
}

// HandleInteraction will echo with a flavour :D
func (h *Handler) HandleInteraction(
	api string,
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

	registry.SendToSameHub(caller.ClientUUID, responseMessage)
}

func (h *Handler) ListensTo() []string {
	return []string{name}
}

func (h *Handler) Name() string {
	return name
}

func (h *Handler) NewClient(
	clientID identifier.Client,
	reg registry.Registry,
) {
}

func (h *Handler) ClientQuit(
	clientID identifier.Client,
	reg registry.Registry,
) {
}
