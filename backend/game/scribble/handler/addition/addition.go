package addition

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name string = "addition"
)

type Addition struct{}

type AdditionMessage struct {
	Num1 float64 `json: "num1"`
	Num2 float64 `json: "num2"`
}

type AdditionResponse struct {
	Result float64 `json: "result"`
}

func Get() *Addition {
	return &Addition{}
}

func (a *Addition) HandleInteraction(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var addmsg AdditionMessage
	err := json.Unmarshal(message, &addmsg)
	if err != nil {
		log.Errorf("unable to unmarshal message: %v", err)
		return
	}

	result := addmsg.Num1 + addmsg.Num2

	payload := AdditionResponse{
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
		log.Errorf("unable to marshal the full response: %v", err)
		return
	}

	registry.SendToCaller(caller, responseMessage)
}

func (a *Addition) Name() string {
	return name
}
