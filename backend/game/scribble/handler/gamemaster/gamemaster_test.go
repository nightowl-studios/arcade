package gamemaster

import (
	"encoding/json"
	"fmt"
	"log"
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
		GameMasterAPI: string(WaitForStart),
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
		GameMasterAPI: string(WaitForStart),
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

var TestGetGameInfoProvider = []struct {
	handler                   Handler
	expectedRemainingDuration time.Duration
}{
	{
		handler: Handler{
			clientList: ClientList{
				currentlySelected: 1,
				clients: []client{
					{
						ClientUUIDStruct: identifier.ClientUUIDStruct{"AAA"},
						GuessedRight:     false,
					},
					{
						ClientUUIDStruct: identifier.ClientUUIDStruct{"BBB"},
						GuessedRight:     false,
					},
					{
						ClientUUIDStruct: identifier.ClientUUIDStruct{"CCC"},
						GuessedRight:     true,
					},
				},
				totalScore: map[string]int{
					"AAA": 500,
					"BBB": 500,
					"CCC": 400,
				},
				roundScore: map[string]int{
					"AAA": 0,
					"BBB": 0,
					"CCC": 200,
				},
			},
			gameState:        WordSelect,
			round:            1,
			chosenWord:       "hello",
			hintString:       "_ e _ _ _",
			maxRounds:        3,
			wordChoices:      3,
			selectTopicTimer: time.Second * 10,
			playTimeTimer:    time.Second * 30,
			timerStartTime:   time.Now().Add(time.Second * -5),
		},
		expectedRemainingDuration: time.Second * 5, // selectTopicTimer == 10s - 5s
	},
	{
		handler: Handler{
			clientList: ClientList{
				currentlySelected: 1,
				clients: []client{
					{
						ClientUUIDStruct: identifier.ClientUUIDStruct{"AAA"},
						GuessedRight:     false,
					},
					{
						ClientUUIDStruct: identifier.ClientUUIDStruct{"BBB"},
						GuessedRight:     false,
					},
					{
						ClientUUIDStruct: identifier.ClientUUIDStruct{"CCC"},
						GuessedRight:     true,
					},
				},
				totalScore: map[string]int{
					"AAA": 100,
					"BBB": 200,
					"CCC": 300,
				},
				roundScore: map[string]int{
					"AAA": 10,
					"BBB": 20,
					"CCC": 200,
				},
			},
			gameState:        PlayTime,
			round:            1,
			chosenWord:       "happy",
			hintString:       "_ _ p p _",
			maxRounds:        2,
			wordChoices:      2,
			selectTopicTimer: time.Second * 10,
			playTimeTimer:    time.Second * 30,
			timerStartTime:   time.Now().Add(time.Second * -15),
		},
		expectedRemainingDuration: time.Second * 15, // playTime == 30s - 15s
	},
}

// TestGetGameInfo will test that when a user joins an 'existing' game session,
// they can ask the Gamemaster for all relevant info to support someone joining
// for the first time, or re-joining the game after disconnecting
func TestGetGameInfo(t *testing.T) {
	requestGameMsg := Receive{
		GameMasterAPI: RequestCurrentGameInfo,
	}
	requestGameMsgByte, err := json.Marshal(requestGameMsg)
	if err != nil {
		log.Fatalf("unable to marshal message :%v", err)
	}
	ID := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}

	for testTrial, testVal := range TestGetGameInfoProvider {
		gameMaster := testVal.handler
		testName := fmt.Sprintf("testTrial :%v", testTrial)
		t.Run(testName, func(t *testing.T) {
			var reg mocks.Registry
			reg.On(
				"SendToCaller",
				// ID
				mock.MatchedBy(func(client identifier.ClientUUIDStruct) bool {
					if client.UUID == ID.ClientUUID.UUID {
						return true
					}
					return false
				}),
				// RequestCurrentGameInfoSend bytes
				mock.MatchedBy(func(sentBytes []byte) bool {
					var gameMsg game.Message
					err := json.Unmarshal(sentBytes, &gameMsg)
					if err != nil {
						t.Fatalf("unable to unmarshal send: %v", err)
						return false
					}
					var send Send
					err = json.Unmarshal(gameMsg.Payload, &send)
					if err != nil {
						t.Fatalf("unable to unmarshal: %v", err)
						return false
					}
					sent := send.CurrentGameInfo
					if !reflect.DeepEqual(sent.Clients, gameMaster.clientList.clients) {
						t.Errorf("got: %v, expected: %v", sent.Clients, gameMaster.clientList.clients)
					}
					if sent.GameState != gameMaster.gameState {
						t.Errorf("got: %v, expected: %v", sent.GameState, gameMaster.gameState)
					}
					if sent.Round != gameMaster.round {
						t.Errorf("got: %v, expected: %v", sent.Round, gameMaster.round)
					}
					if sent.HintString != gameMaster.hintString {
						t.Errorf("got: %v, expected: %v", sent.HintString, gameMaster.hintString)
					}
					if sent.MaxRounds != gameMaster.maxRounds {
						t.Errorf("got: %v, expected: %v", sent.MaxRounds, gameMaster.maxRounds)
					}
					durationDiff := testVal.expectedRemainingDuration - sent.TimerRemaining
					if durationDiff > time.Second {
						t.Errorf(
							"got: %v(s), expected: %v(s)",
							sent.TimerRemaining.Seconds(),
							testVal.expectedRemainingDuration.Seconds(),
						)
					}

					return true
				}),
			)

			gameMaster.reg = &reg
			gameMaster.HandleInteraction(name, requestGameMsgByte, ID, &reg)
			reg.AssertExpectations(t)
		})
	}
}

var TestGetRemainingTimeProvider = []struct {
	startTime     time.Time
	now           time.Time
	timerDuration time.Duration
	expectedDur   time.Duration
}{
	{
		startTime:     time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC), // 10s
		now:           time.Date(2020, 10, 10, 10, 10, 30, 10, time.UTC), // 30s - diff of 20s
		timerDuration: time.Second * 60,
		expectedDur:   time.Second * 40, // diff of 20s
	},
	{
		startTime:     time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC),
		now:           time.Date(2020, 10, 10, 10, 11, 9, 10, time.UTC),
		timerDuration: time.Second * 60,
		expectedDur:   time.Second * 1,
	},
	{
		startTime:     time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC),
		now:           time.Date(2020, 10, 10, 10, 11, 10, 10, time.UTC),
		timerDuration: time.Second * 60,
		expectedDur:   time.Second * 0,
	},
	{
		startTime:     time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC),
		now:           time.Date(2020, 10, 10, 10, 11, 11, 10, time.UTC),
		timerDuration: time.Second * 60,
		expectedDur:   time.Second * 0,
	},
}

// TestGetRemainingTime will unit test the function to make sure it outputs
// the right durations
func TestGetRemainingTime(t *testing.T) {
	for testNum, testVal := range TestGetRemainingTimeProvider {
		testName := fmt.Sprintf("%v", testNum)
		t.Run(testName, func(t *testing.T) {

			outDuration := getRemainingTime(
				testVal.startTime,
				testVal.now,
				testVal.timerDuration,
			)

			if outDuration != testVal.expectedDur {
				t.Errorf(
					"got: %v(s), expected: %v(s)",
					outDuration.Seconds(),
					testVal.expectedDur.Seconds(),
				)
			}

		})
	}
}

// TestWordSelectTimeout will test the timeout functionality of WordSelect
func TestWordSelectTimeout(t *testing.T) {
	wordSelectTimer := time.Second * 3 // to make the test not too long
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
	ID1 := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	ID2 := identifier.Client{
		ClientUUID: identifier.ClientUUIDStruct{"AAA"},
		HubName:    identifier.HubNameStruct{"BBB"},
	}
	gameMaster.NewClient(ID1, &reg)
	gameMaster.NewClient(ID2, &reg)

	wordHint.On("GiveHint", mock.Anything).Return("hello")
	wordFactory.On("GenerateWordList", 3).Return(wordChoices)
	reg.On("SendToSameHubExceptCaller", mock.Anything, mock.Anything)
	reg.On("SendToCaller", mock.Anything, mock.Anything)
	reg.On("SendToSameHub", mock.Anything, mock.Anything)

	go gameMaster.run()
	time.Sleep(time.Millisecond * 100)

	if gameMaster.clientList.currentlySelected != 0 {
		t.Errorf(
			"got currentlySelected: %v, expected: %v",
			gameMaster.clientList.currentlySelected,
			0,
		)
	}

	time.Sleep(wordSelectTimer)

	if gameMaster.clientList.currentlySelected != 1 {
		t.Errorf(
			"got currentlySelected: %v, expected: %v",
			gameMaster.clientList.currentlySelected,
			1,
		)
	}

}
