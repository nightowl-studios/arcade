package gamemaster

import (
	"encoding/json"
	"time"

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
	ShowResults
	EndGame
)

var (
	// api's we want to listen to
	names []string = []string{
		"chat",
		"game",
	}
)

type GameMasterSend struct {
	PlayerSend PlayerSelectSend `json:"PlayerSelectSend"`
}

type ClientList struct {
	initialized      bool
	nextToBeSelected int
	clients          []identifier.ClientUUIDStruct
}

type Handler struct {
	reg       registry.Registry
	gameState State
	round     int
	maxRounds int

	selectTopicTimer time.Duration
	clientList       ClientList
}

func Get(reg registry.Registry) *Handler {
	// hardcoded values
	handler := &Handler{
		maxRounds:        3,
		round:            0,
		gameState:        PlayerSelectTopic,
		selectTopicTimer: 5 * time.Second,
	}
	go handler.run()
	return handler
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
	// if a user joins halfway, they'll be appended to the end of the
	// clientList
	h.clientList.clients = append(h.clientList.clients, clientID.ClientUUID)
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

// run is the function that should be called as a thread
// It will handle the state machine which is affected by a timer, and probably
// by the chat input
// notes: Might need to add some sort of channel to get chat input into the run
func (h *Handler) run() {

	for {
		switch h.gameState {
		case PlayerSelectTopic:
			h.playerSelectTopic()
			h.gameState = PlayTime
		case PlayTime:
			h.playTime()
			h.gameState = ScoreTime
		case ScoreTime:
			h.scoreTime()
			// now we increment the round after showing score for
			// the current round
			h.round++
			if h.round >= h.maxRounds {
				h.gameState = ShowResults
			} else {
				h.gameState = PlayerSelectTopic
			}
		case ShowResults:
			h.showResults()
			h.gameState = EndGame
			break
		}

		if h.gameState == EndGame {
			break
		}
	}
}

type PlayerSelectSend struct {
	ChosenUUID string   `json:"chosenUUID"`
	Choices    []string `json:"choices,omitempty"`
}

type PlayerSelectReceive struct {
	Choice int `json:"choice"`
}

func (h *Handler) playerSelectTopic() {

	if h.clientList.initialized == false {
		clients := h.reg.GetClientSlice()
		for _, client := range clients {
			h.clientList.clients = append(
				h.clientList.clients,
				client.ClientUUID,
			)
		}
		h.clientList.initialized = true
		h.clientList.nextToBeSelected = 0
	}

	selectedClient := h.clientList.clients[h.clientList.nextToBeSelected]

}

func (h *Handler) playTime() {

}

func (h *Handler) scoreTime() {

}
func (h *Handler) showResults() {

}
