import { EventBus } from "@/eventBus";
import { Event } from "@/events";
import Player from "./entities/player";
import { ScribbleEvent } from "./scribbleEvent";
import { ChoosingWord, Drawing, GameOver, Guessing, WaitingForPlayerToChooseWord, WaitingInLobby } from "./states/gameStates";

const NANOSECOND_TO_SECONDS_FACTOR = 1000000000;

export default class GameManager {
    constructor(gameController, applicationStoreService, storeService) {
        this.gameController = gameController;
        this.applicationStoreService = applicationStoreService;
        this.storeService = storeService;
    }

    handle(event, data) {
        console.log('Event:', event);
        console.log('Data:', data);
        if (event === ScribbleEvent.NEW_PLAYER_JOIN) {
            this.onNewPlayerJoin();
        } else if (event === ScribbleEvent.INITIALIZATION) {
            this.loadGame(data);
        } else if (event === ScribbleEvent.PLAYER_UPDATE) {
            this.onPlayerUpdate(data);
        } else if (event === ScribbleEvent.GAME_EVENT) {
            this.onGameEvent(data);
        } else if (event === ScribbleEvent.DRAW_EVENT) {
            this.onDrawEvent(data);
        } else if (event === ScribbleEvent.CHAT_EVENT) {
            this.onChatEvent(data);
        }
    }

    onChatEvent(data) {
        if (data.payload.history) {
            EventBus.$emit(Event.CHAT_HISTORY, data.payload.history);
        }
        else if (data.payload.message) {
            EventBus.$emit(Event.CHAT_MESSAGE, data.payload.message);
        }
    }

    onDrawEvent(data) {
        EventBus.$emit(Event.CANVAS_UPDATE, data.payload);
    }

    onNewPlayerJoin() {
        this.gameController.authenticate();
        const nickname = this.applicationStoreService.getNickname() == null ? "RANDOM" : this.applicationStoreService.getNickname()
        this.gameController.changeNickname(nickname);
        this.gameController.initGame();
    }

    loadGame(data) {
        const api = data.api
        const payload = data.payload;

        if (api === "auth") {
            // initialize player
            const player = new Player();
            this.storeService.setPlayer(player);

            const uuid = payload.uuid;
            console.log("Player uuid set to: " + uuid)
            this.storeService.setPlayerUuid(uuid);
        } else if (api === "hub") {
            const players = this.mapToPlayers(payload);
            this.storeService.setPlayers(players);
        } else if (api === "game") {
            // initialize game state
            if (payload.gameMasterAPI === "requestCurrentGameInfo") {
                console.log("Initializing game state...");
                const gameState = payload.requestCurrentGameInfo.gameState;
                this._updateRoundNumber(payload.requestCurrentGameInfo.round);
                if (gameState === "waitForStart") {
                    const state = new WaitingInLobby();
                    console.log("Game state set to: " + WaitingInLobby.STATE);
                    this.storeService.setState(state);
                } else if (gameState === "wordSelect") {
                    const playerDrawingUuid = payload.requestCurrentGameInfo.selectedClient.UUID;
                    const remainingTime = payload.requestCurrentGameInfo.timerRemaining;
                    this._setStateToWaitingForPlayerToChooseWord(playerDrawingUuid, remainingTime);
                } else if (gameState === "playTime") {
                    this._setStateToGuessing(payload.requestCurrentGameInfo.hintString, payload.requestCurrentGameInfo.timerRemaining);
                }

                this.storeService.setLoading(false);
                console.log("Game finished initializing");
            }
        }
    }

    onPlayerUpdate(data) {
        const players = this.mapToPlayers(data.payload);

        const currentPlayerUuid = this.storeService.getPlayerUuid();
        const updatedCurrentPlayer = players.filter(p => p.uuid === currentPlayerUuid)[0];
        this.storeService.setPlayer(updatedCurrentPlayer)

        this.handlePlayersChanged(players);
    }

    onGameEvent(data) {
        const payload = data.payload;
        if (payload.gameMasterAPI === "waitForStart") {
            this.storeService.setPlayerReadyState(payload.waitForStart);
        }
        else if (payload.gameMasterAPI === "wordSelect") {
            const playerUuid = this.storeService.getPlayerUuid();
            if (playerUuid === payload.wordSelect.chosenUUID) {
                this._setStateToChoosingWord(playerUuid, payload.wordSelect.choices, payload.wordSelect.duration);
            } else {
                this._setStateToWaitingForPlayerToChooseWord(playerUuid, payload.wordSelect.duration);
            }
        } else if (payload.gameMasterAPI === "playTime") {
            const currentState = this.storeService.getState();
            if (currentState.state === ChoosingWord.STATE) {
                this._setStateToDrawing(payload.playTimeSend.duration);
            }
            else if (currentState.state === WaitingForPlayerToChooseWord.STATE) {
                this._setStateToGuessing(payload.playTimeSend.hint, payload.playTimeSend.duration);
            } else {
                this._applyScore(payload.playTimeSend.correctClient.UUID, payload.playTimeSend.totalScore)
            }
        }
        else if (payload.gameMasterAPI === "scoreTime") {
            this._updateRoundNumber(payload.scoreTime.round);
        }
        else if (payload.gameMasterAPI === "showResults") {
            const state = new GameOver();
            this.storeService.setState(state);
        }

    }

    _applyScore(correctPlayerUuid, score) {
        const totalScores = score;
        Object.keys(totalScores).forEach((key) => {
            const playerScore = {
                uuid: key,
                score: totalScores[key],
            };
            this.storeService.setPlayerScore(playerScore);
        });

        const player = this.storeService.getPlayerWithUuid(correctPlayerUuid);
        EventBus.$emit(Event.CORRECT_GUESS, player);
    }

    _setStateToChoosingWord(playerUuid, wordChoices, duration) {
        const player = this.storeService.getPlayerWithUuid(playerUuid);
        const state = new ChoosingWord(
            player,
            wordChoices,
            this._convertNanoSecsToSecs(duration)
        );
        console.log("Game state set to: " + ChoosingWord.STATE);
        this.storeService.setState(state);
        EventBus.$emit(Event.TIMER_RESET, state.duration);
    }

    _setStateToWaitingForPlayerToChooseWord(playerUuid, duration) {
        const player = this.storeService.getPlayerWithUuid(playerUuid)
        const state = new WaitingForPlayerToChooseWord(
            player,
            this._convertNanoSecsToSecs(duration)
        );
        console.log("Game state set to: " + WaitingForPlayerToChooseWord.STATE);
        this.storeService.setState(state);
        EventBus.$emit(Event.TIMER_RESET, state.duration);
    }

    _setStateToDrawing(duration) {
        const selectedWord = this.storeService.getWordSelected();
        const state = new Drawing(
            selectedWord,
            this._convertNanoSecsToSecs(duration)
        );
        console.log("Game state set to: " + Drawing.STATE);
        this.storeService.setState(state);
        EventBus.$emit(Event.TIMER_RESET, state.duration);
    }

    _setStateToGuessing(hint, duration) {
        const state = new Guessing(
            hint,
            this._convertNanoSecsToSecs(duration)
        );
        console.log("Game state set to " + Guessing.STATE);
        this.storeService.setState(state);
        EventBus.$emit(Event.TIMER_RESET, state.duration);
    }



    getCurrentState() {
        return this.storeService.getState();
    }

    mapToPlayers(payload) {
        const players = payload.connectedClients.map(
            (client) =>
                new Player(
                    client.clientUUID.UUID,
                    client.nickname,
                    client.joinOrder
                )
        );

        return players;
    }

    handlePlayersChanged(players) {
        const currentPlayers = this.storeService.getPlayers();
        if (currentPlayers.length !== 0) {
            const playersJoined = this.getPlayerListDifference(
                players,
                currentPlayers
            );
            playersJoined.forEach((player) => {
                console.log("Player joined");
                this.storeService.addPlayer(player);
                EventBus.$emit(Event.PLAYER_JOIN, player);
            });

            const playersLeft = this.getPlayerListDifference(
                currentPlayers,
                players
            );
            playersLeft.forEach((player) => {
                console.log("Player left");
                this.storeService.removePlayer(player);
                EventBus.$emit(Event.PLAYER_LEFT, player);
            });

            // Update nickname if changed
            players.forEach((player) => {
                const currentPlayer = this.storeService.getPlayerWithUuid(player.uuid);
                if (currentPlayer.nickname !== player.nickname) {
                    this.storeService.updateNickname(player.uuid, player.nickname);
                }
            })
        } else {
            this.storeService.setPlayers(players)
        }
    }

    getPlayerListDifference(playerList1, playerList2) {
        return playerList1.filter(
            (p) => !playerList2.map((q) => q.uuid).includes(p.uuid)
        );
    }

    _convertNanoSecsToSecs(durationNS) {
        return Math.round(durationNS / NANOSECOND_TO_SECONDS_FACTOR);
    }

    _updateRoundNumber(roundNumber) {
        console.log("updating round number");
        this.storeService.setRoundNumber(roundNumber);
    }
}

