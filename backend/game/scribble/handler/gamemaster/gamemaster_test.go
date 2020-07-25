package gamemaster

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/bseto/arcade/backend/game/scribble/handler/gamemaster/util/point"
	mockWf "github.com/bseto/arcade/backend/mocks/util/wordfactory"
	mocks "github.com/bseto/arcade/backend/mocks/websocket/registry"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/stretchr/testify/mock"
)

func TestWaitForStart(t *testing.T) {
	var reg mocks.Registry
	var wordFactory mockWf.WordFactory
	reg.On("SendToSameHubExceptCaller", mock.Anything, mock.Anything)
	reg.On("SendToCaller", mock.Anything, mock.Anything)
	wordFactory.On("GenerateWordList", 3).Return([]string{"a", "b", "c"})

	gameMaster := Get(&reg)
	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)
	gameMaster.wordFactory = &wordFactory

	if gameMaster.gameState != WaitForStart {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WaitForStart)
	}

	receivedStruct := Receive{
		GameMasterAPI: WaitForStart,
		WaitForStartReceive: WaitForStartReceive{
			StartGame: false,
		},
	}
	b, err := json.Marshal(receivedStruct)
	if err != nil {
		t.Fatalf("unable to marshal struct: %v", err)
	}
	gameMaster.HandleInteraction(gameMaster.Name(), b, ID, &reg)
	if gameMaster.gameState != WaitForStart {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WaitForStart)
	}

	receivedStruct.WaitForStartReceive.StartGame = true
	b, err = json.Marshal(receivedStruct)
	if err != nil {
		t.Fatalf("unable to marshal struct: %v", err)
	}

	gameMaster.HandleInteraction(gameMaster.Name(), b, ID, &reg)
	time.Sleep(time.Millisecond * 100)

	if gameMaster.gameState != WordSelect {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WordSelect)
	}

	reg.AssertExpectations(t)
	wordFactory.AssertExpectations(t)
}

func TestWordSelect(t *testing.T) {
	startingState := WordSelect
	wordSelectTimer := time.Second * 10

	var reg mocks.Registry
	var wordFactory mockWf.WordFactory
	gameMaster := &Handler{
		reg:              &reg,
		maxRounds:        3,
		wordChoices:      3,
		round:            0,
		gameState:        startingState, // start at wordSelect
		selectTopicTimer: wordSelectTimer,
		playTimeTimer:    180 * time.Second,
		playTimeChan:     make(chan PlayTimeChanReceive),
		selectTopicChan:  make(chan WordSelectReceive),
		waitForStartChan: make(chan WaitForStartReceive),
		pointHandler:     point.Get(),
		wordFactory:      &wordFactory,
	}
	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)

	wordFactory.On("GenerateWordList", 3).Return([]string{"a", "b", "c"})
	reg.On("SendToSameHubExceptCaller",
		mock.Anything,
		mock.MatchedBy(func(msg Send) bool {
			return false
		}),
	)

	go gameMaster.run()

	time.Sleep(time.Millisecond * 500)

	reg.AssertExpectations(t)
	wordFactory.AssertExpectations(t)

}
