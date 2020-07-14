package draw

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name = "draw"
)

type Handler struct {
}

type LineCap string

const (
	Round LineCap = "round"
)

type BrushStyle struct {
	BrushSize  int32  `json:"brushSize"`
	BrushColor string `json:"brushColor"`
}

type DrawPosition struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type DrawMessage struct {
	From       DrawPosition `json:"from"`
	To         DrawPosition `json:"to"`
	BrushStyle BrushStyle   `json:"brushStyle"`
	LineCap    LineCap      `json:"lineCap"`
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
	var msg DrawMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal message: %v", err)
		return
	}

	h.forwardDrawMessage(msg, caller, registry)

	return
}

func (h *Handler) forwardDrawMessage(
	drawMessage DrawMessage,
	clientID identifier.Client,
	reg registry.Registry,
) {

	drawBytes, err := game.MessageBuild(name, drawMessage)
	if err != nil {
		log.Errorf("unable to build message: %v", err)
		return
	}

	reg.SendToSameHubExceptCaller(clientID, drawBytes)
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
