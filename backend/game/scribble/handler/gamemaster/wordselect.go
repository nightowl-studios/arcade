package gamemaster

import (
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
)

// WordSelectSend defines the message that the "wordSelect" state
// will send to the front end
type WordSelectSend struct {
	// ChosenUUID is the UUID of the client that was chosen
	// to pick a word
	ChosenUUID string `json:"chosenUUID"`
	// Choices is a slice of strings that the Chosen client is allowed to choose
	// from. This Choices field is only sent to the Chosen client
	Choices []string `json:"choices,omitempty"`

	// Duration is the duration in nanoseconds that the user is allowed to
	// have to choose a word from the 'choices'
	Duration time.Duration `json:"duration"`

	LockCanvas bool `json:"lockCanvas"`

	Round int `json:"round"`
}

type WordSelectReceive struct {
	// If the user chose a word, set this bool to true
	// If the user did not choose a word and timed out, then set this bool to false
	WordChosen bool `json:"wordChosen"`
	// Choice is given back to the gamemaster as an int. If the user selected
	// the first option, give this back in terms of slice (list) indices - So
	// the front end should give 0.
	Choice int `json:"choice"`
}

// wordSelect will in-order select a user from the clientList
// and then provide them with 3 word choices.
// this function will also let the other players know that the selected player
// is currently choosing a word
func (h *Handler) wordSelect() {
	wordChoices := h.wordFactory.GenerateWordList(h.wordChoices)
	selectedClient := h.clientList.clients[h.clientList.currentlySelected]

	selectedPlayerMsg := Send{
		GameMasterAPI: WordSelect,
		WordSelectSend: WordSelectSend{
			LockCanvas: true,
			Duration:   h.selectTopicTimer,
			ChosenUUID: selectedClient.UUID,
			Round:      h.round,
		},
	}
	selectedPlayerBytes, err := game.MessageBuild("game", selectedPlayerMsg)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return
	}
	h.reg.SendToSameHubExceptCaller(selectedClient.ClientUUIDStruct, selectedPlayerBytes)

	// We did not want to send the other players the wordChoices just in case
	// they're (zachary) snooping the websocket messages :P
	selectedPlayerMsg.WordSelectSend.Choices = wordChoices
	selectedPlayerMsg.WordSelectSend.LockCanvas = false
	selectedPlayerBytes, err = game.MessageBuild("game", selectedPlayerMsg)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return
	}
	h.reg.SendToCaller(selectedClient.ClientUUIDStruct, selectedPlayerBytes)

	selectTopicTime := time.NewTimer(h.selectTopicTimer)
	h.timerStartTime = time.Now()
	select {
	case <-selectTopicTime.C:
		h.clientList.currentlySelected++
		h.WrapUserAndRound()
		h.changeGameStateTo(WordSelect)
	case msg := <-h.selectTopicChan:
		if msg.WordChosen == false {
			h.changeGameStateTo(WordSelect)
		} else {
			h.chosenWord = wordChoices[msg.Choice]
			h.changeGameStateTo(PlayTime)
		}
	case <-h.endChan:
		// we need to enter the run() loop so we can exit
		return
	}
	return
}
