package hubapi

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name string = "hub"
)

// HubAPI is the struct that this handler expects as input
type HubAPI struct {
	RequestLobbyDetails bool   `json:"requestLobbyDetails"`
	ChangeNameTo        string `json:"changeNameTo"`
}

// HubAPIReply is the struct that this handler will send to the clients
type HubAPIReply struct {
	ConnectedClients []*identifier.UserDetails `json:"connectedClients,omitempty"`
}

type Handler struct{}

func Get() *Handler {
	return &Handler{}
}

func (r *Handler) ListensTo() []string {
	return []string{name}
}
func (r *Handler) Name() string {
	return name
}

func (h *Handler) SendLobbyDetails(
	clientID identifier.Client,
	reg registry.Registry,
) {
	var hubAPIReply HubAPIReply
	hubAPIReply.ConnectedClients = reg.GetClientSlice()
	b, err := json.Marshal(hubAPIReply)
	if err != nil {
		log.Errorf("unable to marshal response: %v", err)
		return
	}

	reply := game.Message{
		API:     name,
		Payload: b,
	}

	replyBytes, err := json.Marshal(reply)
	if err != nil {
		log.Errorf("unable to marshal response: %v", err)
	}

	reg.SendToSameHub(clientID.ClientUUID, replyBytes)
}

func (h *Handler) HandleInteraction(
	api string,
	message json.RawMessage,
	clientID identifier.Client,
	reg registry.Registry,
) {
	var hubAPI HubAPI
	err := json.Unmarshal(message, &hubAPI)
	if err != nil {
		log.Errorf("unable to unmarshal into hubAPI: %v", err)
		return
	}

	var hubAPIReply HubAPIReply
	sendToCallerOnly := true

	if hubAPI.ChangeNameTo != "" {
		reg.GetClientUserDetail(clientID.ClientUUID).ChangeNickName(hubAPI.ChangeNameTo)
		// we want to respond with the new nickname
		hubAPI.RequestLobbyDetails = true
		// we want to tell everyone
		sendToCallerOnly = false
	}

	if hubAPI.RequestLobbyDetails == true {
		hubAPIReply.ConnectedClients = reg.GetClientSlice()
	}

	b, err := json.Marshal(hubAPIReply)
	if err != nil {
		log.Errorf("unable to marshal response: %v", err)
		return
	}

	reply := game.Message{
		API:     name,
		Payload: b,
	}

	replyBytes, err := json.Marshal(reply)
	if err != nil {
		log.Errorf("unable to marshal response: %v", err)
	}

	if sendToCallerOnly {
		reg.SendToCaller(clientID.ClientUUID, replyBytes)
	} else {
		reg.SendToSameHub(clientID.ClientUUID, replyBytes)
	}

	return
}

func (h *Handler) NewClient(
	clientID identifier.Client,
	reg registry.Registry,
) {
	h.SendLobbyDetails(clientID, reg)
}

func (h *Handler) ClientQuit(
	clientID identifier.Client,
	reg registry.Registry,
) {
	h.SendLobbyDetails(clientID, reg)
}
