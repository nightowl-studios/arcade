package chat

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name = "chat"
)

// RecieveChat is the struct that this handler expects as input
type ReceiveChat struct {
	Message        string `json:"message"`
	RequestHistory bool   `json:"requestHistory"`
}

// ChatReply is the struct that this handler will send to the clients
type ChatReply struct {
	Message ChatMessage  `json:"message"`
	History *ChatHistory `json:"history,omitempty"`
}

type ChatHistory struct {
	History []ChatMessage `json:"history,omitempty"`
}
type ChatMessage struct {
	Timestamp ChatTime                `json:"timestamp"`
	Sender    *identifier.UserDetails `json:"sender"`
	Message   string                  `json:"message"`
}
type Handler struct {
	chatHistoryLock sync.RWMutex
	chatHistory     ChatHistory
}

type ChatTime time.Time

func Get() *Handler {
	return &Handler{}
}

func (c ChatTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(c).Format("2006-01-02T15:04:05Z07:00"))
	return []byte(stamp), nil
}

// HandleInteraction will be given the tools it needs to handle
// any interaction
func (h *Handler) HandleInteraction(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var msg ReceiveChat
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal message: %v", err)
		return
	}

	if msg.RequestHistory == true {
		h.SendHistory(caller, registry)
	} else {
		h.EchoMessage(msg.Message, caller, registry)
	}

	return
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

func (h *Handler) SendHistory(
	clientID identifier.Client,
	reg registry.Registry,
) {
	h.chatHistoryLock.RLock()
	defer h.chatHistoryLock.RUnlock()

	historyBytes, err := game.MessageBuild(name, h.chatHistory)
	if err != nil {
		log.Errorf("unable to build message: %v", err)
		return
	}

	reg.SendToCaller(clientID, historyBytes)
}

func (h *Handler) EchoMessage(
	message string,
	caller identifier.Client,
	registry registry.Registry,
) {
	newChatMessage := ChatMessage{
		Timestamp: ChatTime(time.Now()),
		Sender:    registry.GetClientUserDetail(caller),
		Message:   message,
	}

	chatReply := ChatReply{Message: newChatMessage}

	byteMessage, err := game.MessageBuild(name, chatReply)

	if err != nil {
		log.Errorf("unable to marshal the chat message: %v", err)
		return
	}

	go registry.SendToSameHub(caller, byteMessage)

	h.chatHistoryLock.Lock()
	defer h.chatHistoryLock.Unlock()
	h.chatHistory.History = append(h.chatHistory.History, newChatMessage)

	return
}
