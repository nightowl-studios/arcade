package draw

import (
	"encoding/json"
	"sync"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name = "draw"
)

type ReceiveDraw struct {
	Action         DrawAction `json:"action"`
	RequestHistory bool       `json:"requestHistory"`
}

type DrawReply struct {
	Action  DrawAction   `json:"action"`
	History *DrawHistory `json:"history,omitempty"`
}

type DrawHistory struct {
	History []DrawAction `json:"history,omitempty"`
}

type DrawAction struct {
	From       Point      `json:"from"`
	To         Point      `json:"to"`
	BrushStyle BrushStyle `json:"brushStyle"`
	LineCap    LineCap    `json:"lineCap"`
}

type Point struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type BrushStyle struct {
	BrushSize  int32  `json:"brushSize"`
	BrushColor string `json:"brushColor"`
}

type LineCap string

const (
	Round LineCap = "round"
)

type Handler struct {
	drawHistoryLock sync.RWMutex
	drawHistory     DrawHistory
}

func Get() *Handler {
	return &Handler{}
}

// HandleInteraction will be given the tools it needs to handle
// any interaction
func (h *Handler) HandleInteraction(
	api string,
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var msg ReceiveDraw
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal message: %v", err)
		return
	}

	if msg.RequestHistory == true {
		h.SendHistory(caller, registry)
	} else {
		h.ForwardAction(msg.Action, caller, registry)
	}

	return
}

func (h *Handler) forwardAction(
	drawAction DrawAction,
	clientID identifier.Client,
	reg registry.Registry,
) {

	drawBytes, err := game.MessageBuild(name, drawAction)
	if err != nil {
		log.Errorf("unable to build message: %v", err)
		return
	}

	reg.SendToSameHubExceptCaller(clientID.ClientUUID, drawBytes)
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

func (h *Handler) ListensTo() []string {
	return []string{name}
}

func (h *Handler) Name() string {
	return name
}

func (h *Handler) SendHistory(
	clientID identifier.Client,
	reg registry.Registry,
) {
	h.drawHistoryLock.RLock()
	defer h.drawHistoryLock.RUnlock()

	historyBytes, err := game.MessageBuild(name, h.drawHistory)
	if err != nil {
		log.Errorf("unable to build message: %v", err)
		return
	}

	reg.SendToCaller(clientID.ClientUUID, historyBytes)
}

func (h *Handler) ForwardAction(
	drawAction DrawAction,
	caller identifier.Client,
	registry registry.Registry,
) {
	drawReply := DrawReply{Action: drawAction}

	byteMessage, err := game.MessageBuild(name, drawReply)

	if err != nil {
		log.Errorf("unable to marshal the chat message: %v", err)
		return
	}

	go registry.SendToSameHubExceptCaller(caller.ClientUUID, byteMessage)

	h.drawHistoryLock.Lock()
	defer h.drawHistoryLock.Unlock()
	h.drawHistory.History = append(h.drawHistory.History, drawAction)

	return
}
