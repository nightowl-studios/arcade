package draw

import (
	"encoding/json"
	"time"

	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name = "draw"
)

type Handler struct {
	//chatHistoryLock sync.RWMutex
	//chatHistory     ChatHistory
}

type ChatTime time.Time

func Get() *Handler {
	return &Handler{}
}

// HandleInteraction will be given the tools it needs to handle
// any interaction
func (h *Handler) HandleInteraction(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
}

func (h *Handler) NewClient(
	clientID identifier.Client,
	reg registry.Registry,
) {
	// we don't need to send history on a new connection
	//h.SendHistory(clientID, reg)
}

func (h *Handler) ClientQuit(
	clientID identifier.Client,
	reg registry.Registry,
) {
	// stub
}

// Name needs to return a unique name of this GameHandler
// This return will be used for routing
func (h *Handler) Name() string {
	return name
}
