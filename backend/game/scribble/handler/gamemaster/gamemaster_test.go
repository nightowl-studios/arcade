package gamemaster

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/scribble/handler/gamemaster/util/point"
	mockWf "github.com/bseto/arcade/backend/mocks/util/wordfactory"
	mocks "github.com/bseto/arcade/backend/mocks/websocket/registry"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/stretchr/testify/mock"
)

// TestWaitForStart will test that the gamemaster starts in the WaitForStart state.
// Afterwards, it'll check to make sure that it will not proceed to the next
// state unless the correct message has been passed through.
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

	gameMaster.ClientQuit(ID, &reg)

	reg.AssertExpectations(t)
	wordFactory.AssertExpectations(t)
}

// TestWordSelect tests the gamemaster starting at the WordSelect state.
// This test will check that the messages sent to the front end initially are
// correct.
func TestWordSelect(t *testing.T) {
	wordSelectTimer := time.Second * 10
	wordChoices := []string{"a", "b", "c"}

	var reg mocks.Registry
	var wordFactory mockWf.WordFactory
	gameMaster := &Handler{
		reg:              &reg,
		maxRounds:        3,
		wordChoices:      3,
		round:            0,
		gameState:        WordSelect,
		selectTopicTimer: wordSelectTimer,
		playTimeTimer:    180 * time.Second,
		playTimeChan:     make(chan PlayTimeChanReceive),
		selectTopicChan:  make(chan WordSelectReceive),
		waitForStartChan: make(chan WaitForStartReceive),
		EndChan:          make(chan bool),
		pointHandler:     point.Get(),
		wordFactory:      &wordFactory,
	}
	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)

	wordFactory.On("GenerateWordList", 3).Return(wordChoices)
	reg.On(
		"SendToSameHubExceptCaller",
		mock.Anything,
		mock.MatchedBy(func(msg []byte) bool {
			var gameMsg game.Message
			err := json.Unmarshal(msg, &gameMsg)
			if err != nil {
				t.Errorf("unable to unmarshal send: %v", err)
				return false
			}
			if gameMsg.API != name {
				t.Errorf("got: %v, expected: %v", gameMsg.API, name)
			}
			var send Send
			err = json.Unmarshal(gameMsg.Payload, &send)
			if err != nil {
				t.Errorf("unable to unmarshal send: %v", err)
			}
			if send.GameMasterAPI != WordSelect {
				t.Errorf("got: %v, expected: %v", send.GameMasterAPI, WordSelect)
			}
			if len(send.WordSelectSend.Choices) != 0 {
				t.Errorf(
					"there should be no word choices sent to other clients: len == %v",
					len(send.WordSelectSend.Choices),
				)
			}

			return true
		}),
	)
	reg.On("SendToCaller", mock.Anything, mock.MatchedBy(func(msg []byte) bool {
		var gameMsg game.Message
		err := json.Unmarshal(msg, &gameMsg)
		if err != nil {
			t.Errorf("unable to unmarshal send: %v", err)
			return false
		}
		if gameMsg.API != name {
			t.Errorf("got: %v, expected: %v", gameMsg.API, name)
		}
		var send Send
		err = json.Unmarshal(gameMsg.Payload, &send)
		if err != nil {
			t.Errorf("unable to unmarshal send: %v", err)
		}
		if !reflect.DeepEqual(send.WordSelectSend.Choices, wordChoices) {
			t.Errorf("got: %v, expected: %v", send.WordSelectSend.Choices, wordChoices)
		}
		return true
	}))

	go gameMaster.run()

	time.Sleep(time.Millisecond * 100)

	gameMaster.ClientQuit(ID, &reg)

	reg.AssertExpectations(t)
	wordFactory.AssertExpectations(t)
}

// TestWordSelect tests the gamemaster starting at the WordSelect state.
// This test will check that the messages sent to the front end initially are
// correct.
func TestAllClientsQuit(t *testing.T) {
	var reg mocks.Registry
	var wordFactory mockWf.WordFactory
	gameMaster := &Handler{
		reg:              &reg,
		maxRounds:        3,
		wordChoices:      3,
		round:            0,
		gameState:        WaitForStart,
		selectTopicTimer: time.Second * 10,
		playTimeTimer:    180 * time.Second,
		playTimeChan:     make(chan PlayTimeChanReceive),
		selectTopicChan:  make(chan WordSelectReceive),
		waitForStartChan: make(chan WaitForStartReceive),
		EndChan:          make(chan bool),
		pointHandler:     point.Get(),
		wordFactory:      &wordFactory,
	}
	go gameMaster.run()
	time.Sleep(time.Millisecond * 100)

	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	ID2 := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"BBB"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)
	gameMaster.NewClient(ID2, &reg)

	if gameMaster.gameState != WaitForStart {
		t.Errorf(
			"unexpected gamestate: %v, expected: %v",
			gameMaster.gameState,
			WaitForStart,
		)
	}

	gameMaster.ClientQuit(ID, &reg)
	gameMaster.ClientQuit(ID2, &reg)

	time.Sleep(time.Millisecond * 100)

	if gameMaster.gameState != Ended {
		t.Errorf(
			"unexpected gamestate: %v, expected: %v",
			gameMaster.gameState,
			Ended,
		)
	}

}
