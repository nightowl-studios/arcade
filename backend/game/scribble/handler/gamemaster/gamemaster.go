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
	ShowResults
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
	Round     int
	MaxRounds int
}

func Get() *Handler {
	handler := &Handler{
		MaxRounds: 3,
		Round:     0,
		GameState: PlayerSelectTopic,
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

// run is the function that should be called as a thread
// It will handle the state machine which is affected by a timer, and probably
// by the chat input
// notes: Might need to add some sort of channel to get chat input into the run
func (h *Handler) run() {

	for {
		switch h.GameState {
		case PlayerSelectTopic:
			h.playerSelectTopic()
			h.GameState = PlayTime
		case PlayTime:
			h.playTime()
			h.GameState = ScoreTime
		case ScoreTime:
			h.scoreTime()
			// now we increment the round after showing score for
			// the current round
			h.Round++
			if h.Round == h.MaxRounds {
				h.GameState = ShowResults
			} else {
				h.GameState = PlayerSelectTopic
			}
		case ShowResults:
			h.showResults()
			break
		}
	}
}

func (h *Handler) playerSelectTopic() {

}
func (h *Handler) playTime() {

}

func (h *Handler) scoreTime() {

}
func (h *Handler) showResults() {

}
