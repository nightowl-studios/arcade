package gamemaster

import (
	"encoding/json"
	"testing"
	"time"

	mocks "github.com/bseto/arcade/backend/mocks/websocket/registry"
	"github.com/bseto/arcade/backend/websocket/identifier"
)

func TestWaitForStart(t *testing.T) {
	var reg mocks.Registry
	gameMaster := Get(&reg)
	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)

	if gameMaster.gameState != WaitForStart {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WaitForStart)
	}

	receivedStruct := WaitForStartReceive{
		StartGame: false,
	}
	b, err := json.Marshal(receivedStruct)
	if err != nil {
		t.Fatalf("unable to marshal struct: %v", err)
	}
	gameMaster.HandleInteraction(gameMaster.Name(), b, ID, &reg)
	if gameMaster.gameState != WaitForStart {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WaitForStart)
	}

	receivedStruct.StartGame = true
	b, err = json.Marshal(receivedStruct)
	if err != nil {
		t.Fatalf("unable to marshal struct: %v", err)
	}
	gameMaster.HandleInteraction(gameMaster.Name(), b, ID, &reg)
	time.Sleep(time.Millisecond * 100)

	if gameMaster.gameState != WordSelect {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WordSelect)
	}
}
