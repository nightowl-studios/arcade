package checkAnswer

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

// Trying to make a function that handles user's answers and compare to the
// real answer.
const (
	name string = "guess"
)

type CheckAnswer struct{}

type GuessMessage struct {
	Answer string `json:"answer"`
	Guess  string `json:"userGuess"`
}

type GuessResult struct {
	Result string `json:"result"`
}

func Get() *CheckAnswer {
	return &CheckAnswer{}
}

func (i *CheckAnswer) HandleInteraction(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var isCorrect bool
	isCorrect = false
	var msg GuessMessage

	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unable to unmarshal message: %v", err)
		return
	}

	var result string

	if msg.Guess == msg.Answer {
		isCorrect = true
		result = "Your guess is correct!"
	} else {
		result = "Your guess is incorrect."
	}

	payload := GuessResult{
		Result: result,
	}

	payloadbytes, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("unable to marshal message: %v", err)
		return
	}

	response := game.Message{
		API:     name,
		Payload: payloadbytes,
	}

	responseMessage, err := json.Marshal(response)
	if err != nil {
		log.Errorf("unable to marshal message: %v", err)
		return
	}

	registry.SendToCaller(caller, responseMessage)
}

func (i *CheckAnswer) Name() string {
	return name
}
