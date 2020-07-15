package gamemaster

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

type State uint32

// States
const (
	PlayerSelectTopic State = iota
	PlayTime
	ScoreTime
	Results
)

var (
	// api's we want to listen to
	names []string = []string{
		"chat",
		"game",
	}
)

type Handler struct {
	GameState State
}

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
	log.Infof("hello")
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
func (h *Handler) Names() []string {
	return names
}
