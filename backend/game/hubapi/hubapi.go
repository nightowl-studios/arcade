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

type Handler struct{}

func GetHubHandler() *Handler {
	return &Handler{}
}

func (r *Handler) RouterName() string {
	return name
}

type HubAPI struct {
	RequestLobbyDetails bool   `json:"requestLobbyDetails"`
	ChangeNameTo        string `json:"changeNameTo"`
}

type HubAPIReply struct {
	ConnectedClients []*identifier.UserDetails `json:"connectedClients,omitempty"`
}

func (r *Handler) SendLobbyDetails(
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

	reg.SendToSameHub(clientID, replyBytes)
}

func (r *Handler) RouteMessage(
	messageType int,
	message []byte,
	clientID identifier.Client,
	messageErr error,
	reg registry.Registry,
) (isHubRequest bool) {

	/////////////////////
	// CODE TO READ STUFF
	var msg game.Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal into game.Message: %v", err)
		return
	}

	if msg.API == name {
		isHubRequest = true
	} else {
		return
	}

	var hubAPI HubAPI
	err = json.Unmarshal(msg.Payload, &hubAPI)
	if err != nil {
		log.Errorf("unable to unmarshal into hubAPI: %v", err)
		return
	}
	// CODE TO READ STUFF
	/////////////////////

	var hubAPIReply HubAPIReply
	sendToCallerOnly := true

	if hubAPI.ChangeNameTo != "" {
		reg.GetClientUserDetail(clientID).ChangeNickName(hubAPI.ChangeNameTo)
		// we want to respond with the new nickname
		hubAPI.RequestLobbyDetails = true
		// we want to tell everyone
		sendToCallerOnly = false
	}

	if hubAPI.RequestLobbyDetails == true {
		hubAPIReply.ConnectedClients = reg.GetClientSlice()
	}

	/////////////////////
	// CODE TO GENERATE RESPONSE BELOW HERE
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
		reg.SendToCaller(clientID, replyBytes)
	} else {
		reg.SendToSameHub(clientID, replyBytes)
	}
	// CODE TO GENERATE RESPONSE BELOW HERE
	/////////////////////

	return
}
