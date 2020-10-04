package gamemaster

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

type PlayTimeSend struct {
	Hint          string                      `json:"hint,omitempty"`
	Duration      time.Duration               `json:"duration,omitempty"`
	TotalScore    map[string]int              `json:"totalScore,omitempty"`
	RoundScore    map[string]int              `json:"roundScore,omitempty"`
	CorrectClient identifier.ClientUUIDStruct `json:"correctClient,omitempty"`
}

type PlayTimeReceive struct {
	Message string `json:"message"`
}

type PlayTimeChanReceive struct {
	Timeout    bool `json:"timeout"`
	AllCorrect bool `json:"allCorrect"`
}

func (h *Handler) playTime() {
	// Send the frontend the hint and the duration

	h.hintString = h.wordHint.GiveHint(h.chosenWord)

	send := Send{
		GameMasterAPI: PlayTime,
		PlayTimeSend: PlayTimeSend{
			Hint:     h.hintString,
			Duration: h.playTimeTimer,
		},
	}

	playTimeSendBytes, err := game.MessageBuild(h.Name(), send)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}
	h.reg.SendToSameHub(playTimeSendBytes)

	// stop here until
	// 1) playTime limit up
	// 2) everyone guessed correctly

	// adding 1 for tolerance
	playTime := time.NewTimer(h.playTimeTimer)
	h.timerStartTime = time.Now()
	select {
	case <-playTime.C:
		// exit select and go to changeGameState call
	case msg := <-h.playTimeChan:
		if msg.AllCorrect {
			// we gucci
		}
	case <-h.endChan:
		// we need to enter the run() loop so we can exit
		// don't use break or else we go to ScoreTime
		return
	}
	h.changeGameStateTo(ScoreTime)
	return
}

// snooping in on "chat"
func (h *Handler) handlePlayMessages(
	api string,
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	switch api {
	case "chat":
		h.handlePlayChatMessages(message, caller, registry)
	default:
		log.Errorf("unknown api :'%v' inside gamemaster", api)
	}
}

func (h *Handler) handlePlayChatMessages(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var msg PlayTimeReceive
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Fatalf("unable to unmarshal message: %v", err)
	}
	currSelected := h.clientList.clients[h.clientList.currentlySelected]
	if caller.ClientUUID.UUID == currSelected.ClientUUIDStruct.UUID {
		// do not allow for the person drawing to receive points if they guess
		// the correct word
		return
	}

	equal := strings.EqualFold(msg.Message, h.chosenWord)
	if !equal {
		return
	}
	points := h.pointHandler.GetPoints()
	h.clientList.totalScore[caller.ClientUUID.UUID] += points
	h.clientList.roundScore[caller.ClientUUID.UUID] = points

	send := Send{
		GameMasterAPI: PlayTime,
		PlayTimeSend: PlayTimeSend{
			Hint:          h.hintString,
			TotalScore:    h.clientList.totalScore,
			RoundScore:    h.clientList.roundScore,
			CorrectClient: caller.ClientUUID,
		},
	}

	sendBytes, err := game.MessageBuild(h.Name(), send)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}
	h.reg.SendToSameHub(sendBytes)

	if h.allCorrect(caller) {
		h.resetGuessRight()
		h.playTimeChan <- PlayTimeChanReceive{
			AllCorrect: true,
		}
	}
}

// allCorrect will return true if all the players other than the
// caller has guessed the word correctly
func (h *Handler) allCorrect(caller identifier.Client) bool {
	// if everyone guessed right, then let playTime know
	allCorrect := true
	for index, client := range h.clientList.clients {
		if client.UUID == caller.ClientUUID.UUID {
			h.clientList.clients[index].GuessedRight = true
		}

		if client.UUID == h.clientList.clients[h.clientList.currentlySelected].UUID {
			// we do not want to check if the person drawing guessed it right
			continue
		}

		if h.clientList.clients[index].GuessedRight != true {
			allCorrect = false
			break
		}
	}
	return allCorrect
}

// resetGuessRight will set all the clients `GuessedRight` boolean to false
func (h *Handler) resetGuessRight() {
	for index := range h.clientList.clients {
		h.clientList.clients[index].GuessedRight = false
	}
}
