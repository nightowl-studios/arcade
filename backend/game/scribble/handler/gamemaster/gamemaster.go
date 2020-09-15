// Package gamemaster defines the scribble gamemaster
// The objective of the gamemaster is to control and advance the states
// that the game can have.
//
// For example, in scribble, when the leader of the lobby presses "start",
// all of the other players should also proceed to the scribble game page.
//
// It is up to the gamemaster to define these states and messages. More details
// can be found in the below documentation
package gamemaster

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/scribble/handler/gamemaster/util/point"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/util/wordfactory"
	"github.com/bseto/arcade/backend/util/wordhint"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

// State defines the game state that the game master is currently in
type State string

// States
const (
	// WaitForStart State is a state where the gamemaster is waiting for the
	// leader to click start.
	// Once the leader clicks start, refer to function `waitForStart()` for
	// the behaviour.
	// This state ends when:
	// 1) The leader presses on the start button
	// The next state is:
	// "wordSelect"
	WaitForStart State = "waitForStart"

	// WordSelect State is a state where the gamemaster will - in no particular
	// order, but will be consistent - choose a player and offer the player 3
	// words to choose from. At the same time, the gamemaster will notify all
	// other players which player they are currently waiting on.
	// All canvases should be cleared at this point and if a player is not
	// chosen, their canvas should be locked.
	// This state ends when:
	// 1) The user selects a word from the word choices
	// 2) The user runs out of time
	// The next state is:
	// If 1), the next game state will be "playTime"
	// if 2), the next game state will be "wordSelect"
	WordSelect State = "wordSelect"

	// PlayTime State is a state where the gamemaster will notify all players
	// that the game has started and there should be the "hint" displayed
	// at the top of the page. During this state, the gamemaster will listen
	// in on the chat and determine which players have guessed the right words
	// and reward them appropriately with points. If any user guesses the
	// right words, all users will be notified which user guessed right, and
	// the points they were rewarded.
	// This state ends when:
	// 1) All users have guessed the correct word
	// 2) The timer runs out
	// The next state:
	// regardless of 1) or 2) will be "scoreTime"
	PlayTime State = "playTime"

	// ScoreTime State is a state where the gamemaster will notify all players
	// that the round has ended and give out the score again + say what round
	// the game is currently on.
	// This state ends when:
	// 1) After a timeout - no message required from frontend for this
	// The next state is:
	// 1) If the rounds < maxRounds, "wordSelect"
	// 2) If the rounds == maxRounds, "showResults"
	ScoreTime State = "scoreTime"

	// ShowResults State is a state where the gamemaster will notify all players
	// to show a fancy results page. This page can be anything we want
	// This state ends when:
	// 1) The leader clicks exit
	// The next state is:
	// "endGame"
	ShowResults State = "showResults"

	// EndGame State is a state where the gamemaster will notify all players that
	// the game has officially ended. The frontend can use this event however
	// they'd like. Maybe redirect to a new lobby, or back to the landing page.
	EndGame State = "endGame"

	// Ended state is a state that can be used to check if the run() loop has
	// been stopped.
	Ended State = "ended"
)

// Commands that can be used on the gamemaster
const (
	// RequestCurrentGameInfo can be sent into the `GameMasterAPI`, and the
	// gamemaster will return an assortment of game info
	RequestCurrentGameInfo = "requestCurrentGameInfo"
)

var (
	// api's we want to listen to
	listensTo []string = []string{
		"chat",
		"game",
	}

	name = "game"
)

// Send is a struct that defines what the gamemaster can send to the frontend
// All possible messages (from every state) is defined in this Send struct.
//
// omitempty will make sure that states that are not sending things, won't have
// some empty field in the Send struct.
//
// To further make things easier, whenever a state sends messages to the front-end
// the GameMasterAPI will be filled out with the State name.
// For example, if "wordSelect" was sending a message with the "wordSelect"
// field, the "gameMasterAPI" field will be "wordSelect"
type Send struct {
	GameMasterAPI    State                      `json:"gameMasterAPI"`
	WordSelectSend   WordSelectSend             `json:"wordSelect,omitempty"`
	PlayTimeSend     PlayTimeSend               `json:"playTimeSend,omitempty"`
	ScoreTimeSend    ScoreTimeSend              `json:"scoreTime,omitempty"`
	CurrentGameInfo  RequestCurrentGameInfoSend `json:"requestCurrentGameInfo,omitempty"`
	WaitForStartSend WaitForStartSend           `json:"waitForStart,omitempty"`
	ShowResults      ShowResultsSend            `json:"showResults,omitempty"`
}

// Receive is a struct that defines what the gamemaster expected to
// receive from the frontend
// All possible messages (from every state) is defined in this Receive struct.
// The "GameMasterAPI" field should be used the same way as the Send struct
type Receive struct {
	GameMasterAPI       string              `json:"gameMasterAPI"`
	WaitForStartReceive WaitForStartReceive `json:"waitForStart"`
	WordSelectReceive   WordSelectReceive   `json:"wordSelect"`
}

type client struct {
	identifier.ClientUUIDStruct
	IsReady      bool `json:"isReady"`
	GuessedRight bool `json:"guessedRight"`
}

// ClientList is a struct used internally to track what users are available
// to select from, and their points
type ClientList struct {
	currentlySelected int
	clients           []client

	// Do not delete from this map even when a player quits. They can reconnect
	totalScore map[string]int
	// This is the scores that each player gained in the round
	roundScore map[string]int
}

type Handler struct {
	// current states
	reg            registry.Registry
	clientList     ClientList
	gameStateLock  sync.RWMutex
	gameState      State
	round          int
	chosenWord     string
	hintString     string
	timerStartTime time.Time // used to track the most recent timer

	// communication
	waitForStartChan (chan WaitForStartReceive)
	selectTopicChan  (chan WordSelectReceive)
	playTimeChan     (chan PlayTimeChanReceive)
	endChan          (chan bool)

	// config-like things
	maxRounds        int
	wordChoices      int
	selectTopicTimer time.Duration
	playTimeTimer    time.Duration

	// util things
	pointHandler point.Handler
	wordFactory  wordfactory.WordFactory
	wordHint     wordhint.WordHint
}

func Get(reg registry.Registry) *Handler {
	// hardcoded values
	handler := &Handler{
		reg:              reg,
		maxRounds:        3,
		wordChoices:      3,
		round:            1,
		gameState:        WaitForStart,
		selectTopicTimer: 100 * time.Second,
		playTimeTimer:    100 * time.Second,
		playTimeChan:     make(chan PlayTimeChanReceive),
		selectTopicChan:  make(chan WordSelectReceive),
		waitForStartChan: make(chan WaitForStartReceive),
		endChan:          make(chan bool),
		pointHandler:     point.Get(),
		wordFactory:      wordfactory.GetWordFactory(),
		wordHint:         wordhint.Get(),
	}
	go handler.run()
	return handler
}

// run is the function that should be called as a thread
// It will handle the state machine
func (h *Handler) run() {
	// don't need to RLock for h.gameState in run() because run() is the only
	// thread that writes to it
	for {
		switch h.gameState {
		case WaitForStart:
			log.Infof("waiting for start....")
			h.waitForStart()
		case WordSelect:
			h.wordSelect()
		case PlayTime:
			h.playTime()
		case ScoreTime:
			h.scoreTime()
		case ShowResults:
			h.showResults()
			h.gameState = EndGame
		case EndGame:
		default:
		}

		if h.gameState == EndGame {
			h.changeGameStateTo(Ended)
			break
		}
	}
}

// HandleInteraction will be given the tools it needs to handle
// any interaction
func (h *Handler) HandleInteraction(
	api string,
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var receive Receive
	err := json.Unmarshal(message, &receive)
	if err != nil {
		log.Fatalf("unable to unmarshal message: %v", err)
	}

	switch receive.GameMasterAPI {
	case RequestCurrentGameInfo:
		h.RequestCurrentGameInfo(caller)
		return
	default:
		// skip and continue
	}

	h.gameStateLock.RLock()
	defer h.gameStateLock.RUnlock()
	switch h.gameState {
	case WaitForStart:
		if api == h.Name() {
			log.Infof("sending thing to start channel")
			receive.WaitForStartReceive.clientUUID = caller.ClientUUID.UUID
			h.waitForStartChan <- receive.WaitForStartReceive
		}
	case WordSelect:
		if api == h.Name() {
			if caller.ClientUUID !=
				h.clientList.clients[h.clientList.currentlySelected].ClientUUIDStruct {
				log.Errorf("client: %v tried to send to gamemaster out of turn", caller)
				return
			}
			h.selectTopicChan <- receive.WordSelectReceive
		}
	case PlayTime:
		h.handlePlayMessages(api, message, caller, registry)

	}
}

func (h *Handler) NewClient(
	clientID identifier.Client,
	reg registry.Registry,
) {

	if h.clientList.roundScore == nil {
		h.clientList.roundScore = make(map[string]int)
	}

	if h.clientList.totalScore == nil {
		h.clientList.totalScore = make(map[string]int)
	}
	// if a user joins halfway, they'll be appended to the end of the
	// clientList
	h.clientList.clients = append(
		h.clientList.clients,
		client{
			ClientUUIDStruct: clientID.ClientUUID,
			IsReady:          false,
			GuessedRight:     false,
		},
	)
}

func (h *Handler) ClientQuit(
	clientID identifier.Client,
	reg registry.Registry,
) {

	if len(h.clientList.clients) == 1 {
		h.changeGameStateTo(EndGame)
		return
	}

	userShifted := false
	for i, client := range h.clientList.clients {
		if clientID.ClientUUID.UUID == client.UUID {
			if i < h.clientList.currentlySelected {
				// userShifted is true if the players who came before
				// the nextToBeSelected left the game
				userShifted = true
			}
			// delete index i from clients while maintaining order
			h.clientList.clients = append(
				h.clientList.clients[:i],
				h.clientList.clients[i+1:]...,
			)
		}
	}

	if len(h.clientList.clients) == h.clientList.currentlySelected {
		// if the last user left (len clients == nextToBeSelected only happens
		// in this case) then we want to wrap the order and potentially increment
		// the round
		h.WrapUserAndRound()
	} else if userShifted {
		// we have to shift back down to stay with the same person
		h.clientList.currentlySelected--
	} else {
		// the person who left was after the selected user so no shift needs to
		// happen
	}
}

func (h *Handler) ListensTo() []string {
	return listensTo
}

func (h *Handler) Name() string {
	return name
}

type RequestCurrentGameInfoSend struct {
	Clients        []client      `json:"clients"`
	GameState      State         `json:"gameState"`
	Round          int           `json:"round"`
	HintString     string        `json:"hintString"`
	MaxRounds      int           `json:"maxRounds"`
	TimerRemaining time.Duration `json:"timerRemaining"`
	SelectedClient client        `json:"selectedClient"`
}

func (h *Handler) RequestCurrentGameInfo(
	caller identifier.Client,
) {

	var remainingTime time.Duration
	h.gameStateLock.RLock()
	defer h.gameStateLock.RUnlock()
	switch h.gameState {
	case WordSelect:
		remainingTime = getRemainingTime(
			h.timerStartTime,
			time.Now(),
			h.selectTopicTimer,
		)
	case PlayTime:
		remainingTime = getRemainingTime(
			h.timerStartTime,
			time.Now(),
			h.playTimeTimer,
		)
	}

	send := Send{
		GameMasterAPI: RequestCurrentGameInfo,
		CurrentGameInfo: RequestCurrentGameInfoSend{
			Clients:        h.clientList.clients,
			GameState:      h.gameState,
			Round:          h.round,
			HintString:     h.hintString,
			MaxRounds:      h.maxRounds,
			TimerRemaining: remainingTime,
			SelectedClient: h.clientList.clients[h.clientList.currentlySelected],
		},
	}
	selectedPlayerBytes, err := game.MessageBuild(h.Name(), send)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return
	}
	h.reg.SendToCaller(caller.ClientUUID, selectedPlayerBytes)
}

// getRemainingTime will calculate how much time is remaining on the timer
func getRemainingTime(
	startTime time.Time,
	now time.Time,
	timerDuration time.Duration,
) time.Duration {

	timeElapsed := now.Sub(startTime)
	remainingTime := timerDuration - timeElapsed

	if remainingTime < 0 {
		remainingTime = 0
	}

	return remainingTime
}

func (h *Handler) changeGameStateTo(state State) {
	h.gameStateLock.Lock()
	defer h.gameStateLock.Unlock()

	h.gameState = state
	if state == EndGame {
		h.endChan <- true
	}
}

// WrapUserAndRound will check if the nextToBeSelected is valid.
// If the nextToBeSelected is greater than the length of clients, it will
// wrap the users and increment the round
func (h *Handler) WrapUserAndRound() {
	if len(h.clientList.clients) <= h.clientList.currentlySelected {
		h.clientList.currentlySelected = 0
		h.round++
	}
}
