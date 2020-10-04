package gamemaster

import (
	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
)

type WaitForStartSend struct {
	ClientUUID string `json:"clientUUID"`
	IsReady    bool   `json:"isReady"`
}

type WaitForStartReceive struct {
	clientUUID string
	IsReady    bool `json:"isReady"`
}

// waitForStart will wait for the leader to press the start button.
// when the leader has pressed the start button, there should be a
// incoming message on the h.waitForStartChan in which we can
// continue onto the next gamestate
func (h *Handler) waitForStart() {
	for {
		select {
		case msg := <-h.waitForStartChan:
			for idx, client := range h.clientList.clients {
				if client.UUID == msg.clientUUID {
					h.clientList.clients[idx].IsReady = msg.IsReady
					playerReadyStateChangedMsg := Send{
						GameMasterAPI: WaitForStart,
						WaitForStartSend: WaitForStartSend{
							ClientUUID: msg.clientUUID,
							IsReady:    msg.IsReady,
						},
					}
					playerReadyStateChangedBytes, err :=
						game.MessageBuild("game", playerReadyStateChangedMsg)
					if err != nil {
						log.Fatalf("unable to marshal: %v", err)
						return
					}
					h.reg.SendToSameHub(
						playerReadyStateChangedBytes)
				}
			}
		case <-h.endChan:
			// we need to enter the run() loop so we can exit
			return
		}

		allReady := true
		for _, client := range h.clientList.clients {
			if !client.IsReady {
				allReady = false
				break
			}
		}

		if allReady {
			h.changeGameStateTo(WordSelect)
			break
		}
	}
}
