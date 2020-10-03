package gamemaster

import (
	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
)

type ShowResultsSend struct {
	ShowResults bool `json:"showResults"`
}

func (h *Handler) showResults() {

	showResultsMsg := Send{
		GameMasterAPI: ShowResults,
		ShowResults: ShowResultsSend{
			ShowResults: true,
		},
	}

	scoreTimeBytes, err := game.MessageBuild("game", showResultsMsg)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return
	}
	h.reg.SendToSameHub(scoreTimeBytes)
}
