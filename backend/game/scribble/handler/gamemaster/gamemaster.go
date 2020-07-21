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
	"path/filepath"
	"sync"
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/util/wordfactory"
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
	// "playerSelect"
	WaitForStart State = "waitForStart"

	// PlayerSelect State is a state where the gamemaster will - in no particular
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
	// if 2), the next game state will be "playerSelect"
	PlayerSelect State = "playerSelect"

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
	// 1) If the rounds < maxRounds, "playerSelect"
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
// For example, if "playerSelect" was sending a message with the "playerSelect"
// field, the "gameMasterAPI" field will be "playerSelect"
type Send struct {
	GameMasterAPI    State            `json:"gameMasterAPI"`
	PlayerSelectSend PlayerSelectSend `json:"playerSelect,omitempty"`
	ScoreTimeSend    ScoreTimeSend    `json:"scoreTime,omitempty"`
}

// Receive is a struct that defines what the gamemaster expected to
// receive from the frontend
// All possible messages (from every state) is defined in this Receive struct.
// The "GameMasterAPI" field should be used the same way as the Send struct
type Receive struct {
	GameMasterAPI       State               `json:"gameMasterAPI"`
	WaitForStartRecieve WaitForStartRecieve `json:"waitForStart"`
	PlayerSelectRecieve PlayerSelectReceive `json:"playerSelect"`
}

// ClientList is a struct used internally to track what users are available
// to select from, and their points
type ClientList struct {
	nextToBeSelected     int
	clients              []identifier.ClientUUIDStruct
	clientCorrectGuesses map[identifier.ClientUUIDStruct]bool
}

type Handler struct {
	reg        registry.Registry
	clientList ClientList

	gameStateLock sync.RWMutex
	gameState     State
	round         int
	chosenWord    string

	waitForStartChan (chan WaitForStartRecieve)
	selectTopicChan  (chan PlayerSelectReceive)
	playTimeChan     (chan PlayTimeReceive)

	// config-like things
	maxRounds        int
	wordChoices      int
	selectTopicTimer time.Duration
	playTimeTimer    time.Duration
}

func Get(reg registry.Registry) *Handler {
	// hardcoded values
	handler := &Handler{
		reg:              reg,
		maxRounds:        3,
		wordChoices:      3,
		round:            0,
		gameState:        WaitForStart,
		selectTopicTimer: 10 * time.Second,
		playTimeTimer:    60 * time.Second,
		playTimeChan:     make(chan PlayTimeReceive),
		selectTopicChan:  make(chan PlayerSelectReceive),
		waitForStartChan: make(chan WaitForStartRecieve),
	}
	go handler.run()
	return handler
}

// HandleInteraction will be given the tools it needs to handle
// any interaction
func (h *Handler) HandleInteraction(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
	var receive Receive
	err := json.Unmarshal(message, &receive)
	if err != nil {
		log.Fatalf("unable to unmarshal message: %v", err)
	}

	h.gameStateLock.RLock()
	defer h.gameStateLock.RUnlock()
	switch h.gameState {
	case WaitForStart:
		log.Infof("sending thing to start channel")
		h.waitForStartChan <- receive.WaitForStartRecieve
	case PlayerSelect:
		if caller.ClientUUID != h.clientList.clients[h.clientList.nextToBeSelected] {
			log.Errorf("client: %v tried to send to gamemaster out of turn", caller)
			return
		}
		h.selectTopicChan <- receive.PlayerSelectRecieve
	}

}

func (h *Handler) NewClient(
	clientID identifier.Client,
	reg registry.Registry,
) {
	// if a user joins halfway, they'll be appended to the end of the
	// clientList
	h.clientList.clients = append(h.clientList.clients, clientID.ClientUUID)
}

func (h *Handler) ClientQuit(
	clientID identifier.Client,
	reg registry.Registry,
) {
	userShifted := false
	for i, client := range h.clientList.clients {
		if clientID.ClientUUID.UUID == client.UUID {
			if i < h.clientList.nextToBeSelected {
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

	if len(h.clientList.clients) == h.clientList.nextToBeSelected {
		// if the last user left (len clients == nextToBeSelected only happens
		// in this case) then we want to wrap the order and potentially increment
		// the round
		h.WrapUserAndRound()
	} else if userShifted {
		// we have to shift back down to stay with the same person
		h.clientList.nextToBeSelected--
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
		case PlayerSelect:
			h.playerSelectTopic()
		case PlayTime:
			// we are going to stop the run function cause after this
			// nothing is defined
			return
			h.playTime()
		case ScoreTime:
			h.scoreTime()
		case ShowResults:
			h.showResults()
			h.gameState = EndGame
			break
		}

		if h.gameState == EndGame {
			break
		}
	}
}

type WaitForStartRecieve struct {
	StartGame bool `json:"startGame"`
}

// waitForStart will wait for the leader to press the start button.
// when the leader has pressed the start button, there should be a
// incoming message on the h.waitForStartChan in which we can
// continue onto the next gamestate
func (h *Handler) waitForStart() {
	select {
	case msg := <-h.waitForStartChan:
		if msg.StartGame == true {
			h.changeGameStateTo(PlayerSelect)
		}
	}
}

// PlayerSelectSend defines the message that the "playerSelect" state
// will send to the front end
type PlayerSelectSend struct {
	// ChosenUUID is the UUID of the client that was chosen
	// to pick a word
	ChosenUUID string `json:"chosenUUID"`
	// Choices is a slice of strings that the Chosen client is allowed to choose
	// from. This Choices field is only sent to the Chosen client
	Choices []string `json:"choices,omitempty"`

	// Duration is the duration in nanoseconds that the user is allowed to
	// have to choose a word from the 'choices'
	Duration time.Duration `json:"duration"`

	LockCanvas bool `json:"lockCanvas"`
	LockChat   bool `json:"lockChat"`
}

type PlayerSelectReceive struct {
	// If the user chose a word, set this bool to true
	// If the user did not choose a word and timed out, then set this bool to false
	WordChosen bool `json:"wordChosen"`
	// Choice is given back to the gamemaster as an int. If the user selected
	// the first option, give this back in terms of slice (list) indices - So
	// the front end should give 0.
	Choice int `json:"choice"`
}

// playerSelectTopic will in-order select a user from the clientList
// and then provide them with 3 word choices.
// this function will also let the other players know that the selected player
// is currently choosing a word
func (h *Handler) playerSelectTopic() {
	var wordChoices []string
	for i := 0; i < h.wordChoices; i++ {
		word, err := wordfactory.WordGenerator2(
			filepath.Join(wordfactory.Dir, wordfactory.File),
		)
		if err != nil {
			log.Errorf("unable to get a word, trying again: %v", err)
			word, err = wordfactory.WordGenerator(
				filepath.Join(wordfactory.Dir, wordfactory.File),
			)
			if err != nil {
				log.Fatalf("unable to get a word using WordGenerator1: %v", err)
			}
		}
		wordChoices = append(wordChoices, word)
	}

	selectedClient := h.clientList.clients[h.clientList.nextToBeSelected]
	selectedPlayerMsg := Send{
		GameMasterAPI: PlayerSelect,
		PlayerSelectSend: PlayerSelectSend{
			LockCanvas: true,
			LockChat:   false,
			Duration:   h.selectTopicTimer,
			ChosenUUID: selectedClient.UUID,
		},
	}
	selectedPlayerBytes, err := game.MessageBuild("game", selectedPlayerMsg)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return
	}
	h.reg.SendToSameHubExceptCaller(selectedClient, selectedPlayerBytes)

	// We did not want to send the other players the wordChoices just in case
	// they're (zachary) snooping the websocket messages :P
	selectedPlayerMsg.PlayerSelectSend.Choices = wordChoices
	selectedPlayerMsg.PlayerSelectSend.LockCanvas = false
	selectedPlayerMsg.PlayerSelectSend.LockChat = true
	selectedPlayerBytes, err = game.MessageBuild("game", selectedPlayerMsg)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
		return
	}
	h.reg.SendToCaller(selectedClient, selectedPlayerBytes)

	// adding 1 for tolerance
	selectTopicTime := time.NewTimer(h.selectTopicTimer + 1)
	select {
	case <-selectTopicTime.C:
		h.changeGameStateTo(PlayerSelect)
	case msg := <-h.selectTopicChan:
		if msg.WordChosen == false {
			h.changeGameStateTo(PlayerSelect)
		} else {
			h.chosenWord = wordChoices[msg.Choice]
		}
		h.changeGameStateTo(PlayTime)
	}
	return
}

type PlayTimeSend struct {
	Score int `json:"score,omitempty"`
}

type PlayTimeReceive struct {
	AllCorrect bool
}

func (h *Handler) playTime() {
	// Tell nonselected people to lock their canvases
	// tell selected person to lock their chat

	// stop here until
	// 1) playTime limit up
	// 2) everyone guessed correctly

	// adding 1 for tolerance
	playTime := time.NewTimer(h.playTimeTimer + 1)

	select {
	case <-playTime.C:
	case msg := <-h.playTimeChan:
		if msg.AllCorrect {
			// we gucci
		}
	}
	h.changeGameStateTo(ScoreTime)
	return
}

// snooping in on "chat"
func (h *Handler) handlePlayMessages() {
	// if someone is correct, assign score
	// and then message everyone someone guessed right

	// if everyone guessed right, then send to channel
}

type ScoreTimeSend struct {
	Round int `json:"round"`
}

func (h *Handler) scoreTime() {
	h.clientList.nextToBeSelected++
	h.WrapUserAndRound()

	if h.round >= h.maxRounds {
		h.gameState = ShowResults
	} else {
		h.gameState = PlayerSelect
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
	h.reg.SendToSameHub(selectedClient, scoreTimeBytes)
}

func (h *Handler) showResults() {

}

func (h *Handler) changeGameStateTo(state State) {
	h.gameStateLock.Lock()
	defer h.gameStateLock.Unlock()
	h.gameState = state
}

// WrapUserAndRound will check if the nextToBeSelected is valid.
// If the nextToBeSelected is greater than the length of clients, it will
// wrap the users and increment the round
func (h *Handler) WrapUserAndRound() {
	if len(h.clientList.clients) <= h.clientList.nextToBeSelected {
		h.clientList.nextToBeSelected = 0
		h.round++
	}
}
