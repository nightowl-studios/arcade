package gamemaster

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/generic/chat"
	"github.com/bseto/arcade/backend/game/scribble/handler/gamemaster/util/point"
	mockWf "github.com/bseto/arcade/backend/mocks/util/wordfactory"
	mockWh "github.com/bseto/arcade/backend/mocks/util/wordhint"
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
	reg.On("SendToSameHub", mock.Anything, mock.Anything).Times(2)
	wordFactory.On("GenerateWordList", 3).Return([]string{"a", "b", "c"})

	gameMaster := Get(&reg)
	ID1 := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"ZZZ"},
	}
	ID2 := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"BBB"},
		HubName:    identifier.HubNameStruct{"ZZZ"},
	}
	gameMaster.NewClient(ID1, &reg)
	gameMaster.NewClient(ID2, &reg)
	gameMaster.wordFactory = &wordFactory

	if gameMaster.gameState != WaitForStart {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WaitForStart)
	}

	receivedStruct := Receive{
		GameMasterAPI: WaitForStart,
		WaitForStartReceive: WaitForStartReceive{
			IsReady: true,
		},
	}

	b, err := json.Marshal(receivedStruct)
	if err != nil {
		t.Fatalf("unable to marshal struct: %v", err)
	}

	gameMaster.HandleInteraction(gameMaster.Name(), b, ID1, &reg)
	if gameMaster.gameState != WaitForStart {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WaitForStart)
	}

	gameMaster.HandleInteraction(gameMaster.Name(), b, ID2, &reg)
	time.Sleep(time.Millisecond * 100)
	if gameMaster.gameState != WordSelect {
		t.Errorf("got state: %v, expected: %v", gameMaster.gameState, WordSelect)
	}

	gameMaster.ClientQuit(ID1, &reg)
	gameMaster.ClientQuit(ID2, &reg)
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
	var wordHint mockWh.WordHint

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
		endChan:          make(chan bool),
		pointHandler:     point.Get(),
		wordFactory:      &wordFactory,
		wordHint:         &wordHint,
	}
	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)

	wordHint.On("GiveHint", mock.Anything).Return("hello")
	wordFactory.On("GenerateWordList", 3).Return(wordChoices)
	// Checks that all other clients are sent the message that a user
	// is choosing a word. However, the word itself should not be sent
	// to the other users
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
	// checks that the user that is selecting the word is sent the word
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
	reg.On("SendToSameHub", mock.Anything, mock.Anything).Return(nil)
	go gameMaster.run()
	time.Sleep(time.Millisecond * 100)

	// Now check that if a user selects an option, that the selected option is
	// made into the gamemaster
	receivedStruct := Receive{
		GameMasterAPI: WaitForStart,
		WordSelectReceive: WordSelectReceive{
			WordChosen: true,
			Choice:     1,
		},
	}
	b, err := json.Marshal(receivedStruct)
	if err != nil {
		t.Fatalf("unable to marshal struct: %v", err)
	}
	gameMaster.HandleInteraction(gameMaster.Name(), b, ID, &reg)
	time.Sleep(time.Millisecond * 100)

	if gameMaster.chosenWord != wordChoices[1] {
		t.Errorf(
			"gameMaster word: %v, expected word: %v",
			gameMaster.chosenWord,
			wordChoices[1],
		)
	}

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
		endChan:          make(chan bool),
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

func TestPlayTime(t *testing.T) {
	var reg mocks.Registry
	var wordFactory mockWf.WordFactory
	var wordHint mockWh.WordHint

	chosenWord := "HelloWorld"
	wordHintString := "_ e _ _ _ _ _ _ _ _"

	gameMaster := &Handler{
		reg:              &reg,
		maxRounds:        3,
		wordChoices:      3,
		round:            0,
		gameState:        PlayTime,
		selectTopicTimer: time.Second * 10,
		playTimeTimer:    180 * time.Second,
		playTimeChan:     make(chan PlayTimeChanReceive),
		selectTopicChan:  make(chan WordSelectReceive),
		waitForStartChan: make(chan WaitForStartReceive),
		endChan:          make(chan bool),
		pointHandler:     point.Get(),
		wordFactory:      &wordFactory,
		wordHint:         &wordHint,
		// test specific
		chosenWord: chosenWord,
	}
	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	ID2 := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"BBB"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	ID3 := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"CCC"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)
	gameMaster.NewClient(ID2, &reg)
	gameMaster.NewClient(ID3, &reg)
	wordFactory.On("GenerateWordList", 3).Return([]string{"a", "b", "c"})
	wordHint.On("GiveHint", mock.Anything).Return(wordHintString)
	reg.On("SendToSameHub", mock.Anything, mock.MatchedBy(
		// The message that everyone gets with word hint
		func(b []byte) bool {
			var gameMsg game.Message
			err := json.Unmarshal(b, &gameMsg)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			var send Send
			err = json.Unmarshal(gameMsg.Payload, &send)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			if send.GameMasterAPI != PlayTime {
				return false
			}
			if send.PlayTimeSend.Hint != wordHintString {
				t.Errorf(
					"got wordhint: %v, expected: %v",
					send.PlayTimeSend.Hint,
					wordHintString,
				)
				return false
			}
			if send.PlayTimeSend.CorrectClient.UUID != "" {
				return false
			}
			return true
		},
	))
	reg.On("SendToSameHub", ID2.ClientUUID, mock.MatchedBy(
		// The message that gets send because ID2 send the right word
		func(b []byte) bool {
			var gameMsg game.Message
			err := json.Unmarshal(b, &gameMsg)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			var send Send
			err = json.Unmarshal(gameMsg.Payload, &send)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			if send.GameMasterAPI != PlayTime {
				return false
			}
			if send.PlayTimeSend.CorrectClient.UUID != ID2.ClientUUID.UUID {
				return false
			}
			return true
		},
	))
	reg.On("SendToSameHub", ID3.ClientUUID, mock.MatchedBy(
		// The message that gets send because ID3 send the right word
		func(b []byte) bool {
			var gameMsg game.Message
			err := json.Unmarshal(b, &gameMsg)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			var send Send
			err = json.Unmarshal(gameMsg.Payload, &send)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			if send.GameMasterAPI != PlayTime {
				return false
			}
			if send.PlayTimeSend.CorrectClient.UUID != ID3.ClientUUID.UUID {
				return false
			}
			return true
		},
	))
	reg.On("SendToSameHub", mock.Anything, mock.MatchedBy(
		// The message is sent because state is now scoreTime
		func(b []byte) bool {
			var gameMsg game.Message
			err := json.Unmarshal(b, &gameMsg)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			var send Send
			err = json.Unmarshal(gameMsg.Payload, &send)
			if err != nil {
				t.Errorf("unable to unmarshal :%v", err)
			}
			if send.GameMasterAPI != ScoreTime {
				return false
			}
			if send.ScoreTimeSend.Round != 0 {
				t.Errorf("got round: %v, expected: %v", send.ScoreTimeSend.Round, 1)
			}
			return true
		},
	))
	reg.On("SendToSameHubExceptCaller", mock.Anything, mock.Anything).Return(nil)
	reg.On("SendToCaller", mock.Anything, mock.Anything).Return(nil)

	sendChat := chat.ReceiveChat{
		Message: "helloworld",
	}
	b, err := json.Marshal(sendChat)
	if err != nil {
		t.Errorf("unable to marshal sendChat: %v", err)
	}

	go gameMaster.run()
	time.Sleep(time.Millisecond * 100)

	gameMaster.HandleInteraction("chat", b, ID2, &reg)
	time.Sleep(time.Millisecond * 100)

	gameMaster.HandleInteraction("chat", b, ID3, &reg)
	time.Sleep(time.Millisecond * 100)

	if gameMaster.gameState != WordSelect {
		t.Errorf("gameState: %v, expected: %v", gameMaster.gameState, WordSelect)
	}
	if gameMaster.clientList.currentlySelected != 1 {
		t.Errorf(
			"currentlySelected: %v, expected: %v",
			gameMaster.clientList.currentlySelected,
			1,
		)
	}

	reg.AssertExpectations(t)
	wordHint.AssertExpectations(t)
	wordFactory.AssertExpectations(t)

}

// TestGetGameInfo will test that when a user joins an 'existing' game session,
// they can ask the Gamemaster for all relevant info to support someone joining
// for the first time, or re-joining the game after disconnecting
func TestGetGameInfo(t *testing.T) {
	var reg mocks.Registry
	var wordFactory mockWf.WordFactory
	var wordHint mockWh.WordHint

	currentRound := 1
	chosenWord := "hello"
	hintString := "_ e _ _ _"

	gameMaster := &Handler{
		reg:        &reg,
		gameState:  WordSelect,
		round:      currentRound,
		chosenWord: chosenWord,
		hintString: hintString,

		playTimeChan:     make(chan PlayTimeChanReceive),
		selectTopicChan:  make(chan WordSelectReceive),
		waitForStartChan: make(chan WaitForStartReceive),
		endChan:          make(chan bool),

		maxRounds:        3,
		wordChoices:      3,
		playTimeTimer:    time.Second * 60,
		selectTopicTimer: time.Second * 10,

		pointHandler: point.Get(),
		wordFactory:  &wordFactory,
		wordHint:     &wordHint,
	}

	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID, &reg)
}
