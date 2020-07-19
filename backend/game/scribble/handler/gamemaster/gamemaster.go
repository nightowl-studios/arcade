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
	WaitForStart State = "waitForStart"
	PlayerSelect State = "playerSelect"
	PlayTime     State = "playTime"
	ScoreTime    State = "scoreTime"
	ShowResults  State = "showResults"
	EndGame      State = "endGame"
)

var (
	// api's we want to listen to
	names []string = []string{
		"chat",
		"game",
	}
)

type GameMasterSend struct {
	GameMasterAPI    State            `json:"gameMasterAPI"`
	PlayerSelectSend PlayerSelectSend `json:"playerSelect,omitempty"`
}

type GameMasterReceive struct {
	GameMasterAPI       State               `json:"gameMasterAPI"`
	WaitForStartRecieve WaitForStartRecieve `json:"waitForStart"`
	PlayerSelectRecieve PlayerSelectReceive `json:"playerSelect"`
}

type ClientList struct {
	initialized          bool
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
	var receive GameMasterReceive
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
	// stub
}

// Name needs to return a unique name of this GameHandler
// This return will be used for routing
func (h *Handler) Names() []string {
	return names
}

// run is the function that should be called as a thread
// It will handle the state machine which is affected by a timer, and probably
// by the chat input
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
			h.changeGameStateTo(ScoreTime)
		case ScoreTime:
			h.scoreTime()
			// now we increment the round after showing score for
			// the current round
			h.round++
			if h.round >= h.maxRounds {
				h.gameState = ShowResults
			} else {
				h.gameState = PlayerSelect
			}
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

func (h *Handler) waitForStart() {
	select {
	case msg := <-h.waitForStartChan:
		if msg.StartGame == true {
			h.changeGameStateTo(PlayerSelect)
		}
	}
}

type PlayerSelectSend struct {
	ChosenUUID string   `json:"chosenUUID"`
	Choices    []string `json:"choices,omitempty"`
}

type PlayerSelectReceive struct {
	WordChosen bool `json:"wordChosen"`
	Choice     int  `json:"choice"`
}

func getClientList(userDetails []*identifier.UserDetails) ClientList {
	var clientList ClientList

	for _, client := range userDetails {
		clientList.clients = append(
			clientList.clients,
			client.ClientUUID,
		)
	}
	clientList.initialized = true
	clientList.nextToBeSelected = 0

	return clientList
}

// playerSelectTopic will in-order select a user from the clientList
// and then provide them with 3 word choices.
// this function will also let the other players know that the selected player
// is currently choosing a word
func (h *Handler) playerSelectTopic() {

	if h.clientList.initialized == false {
		clientSlice := h.reg.GetClientSlice()
		h.clientList = getClientList(clientSlice)
	}

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
	selectedPlayerMsg := GameMasterSend{
		PlayerSelectSend: PlayerSelectSend{
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

	h.clientList.nextToBeSelected++
	return
}

type PlayTimeSend struct {
	LockCanvas bool `json:"lockCanvas,omitempty"`
	LockChat   bool `json:"lockChat,omitempty"`
	Score      int  `json:"score,omitempty"`
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

func (h *Handler) scoreTime() {

}
func (h *Handler) showResults() {

}

func (h *Handler) changeGameStateTo(state State) {
	h.gameStateLock.Lock()
	defer h.gameStateLock.Unlock()
	h.gameState = state
}
