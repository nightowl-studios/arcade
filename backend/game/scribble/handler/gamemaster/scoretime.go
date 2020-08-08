package gamemaster

import (
	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
)

type ScoreTimeSend struct {
	Round int `json:"round"`
}

func (h *Handler) scoreTime() {
	h.clientList.currentlySelected++
	h.WrapUserAndRound()
	h.pointHandler.ResetPoints()

	if h.round > h.maxRounds {
		h.changeGameStateTo(ShowResults)
	} else {
		h.changeGameStateTo(WordSelect)
	}
	selectedClient := h.clientList.clients[0]
	scoreTimeMsg := Send{
		GameMasterAPI: ScoreTime,
		ScoreTimeSend: ScoreTimeSend{
			Round: h.round,
		},
	}
	scoreTimeBytes, err := game.MessageBuild("game", scoreTimeMsg)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return
	}
	h.reg.SendToSameHub(selectedClient.ClientUUIDStruct, scoreTimeBytes)
}
