import { EventBus } from "@/eventBus";
import { Event } from "@/events";
import Player from "./entities/player";
import { ScribbleEvent } from "./scribbleEvent";
import { ChoosingWord, Drawing, Guessing, WaitingForPlayerToChooseWord, WaitingInLobby } from "./states/gameStates";

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
        } else {
            const currentState = this.getCurrentState();
            if (currentState == null) {
                this.initGame(data);
            } else {
                // TODO Refactor this to use vuex instead of an event?
                if (data.api === "draw") {
                    EventBus.$emit(data.api, data.payload);
                }

                this.handleEvent(data);
            }
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
                if (gameState === "waitForStart") {
                    const state = new WaitingInLobby();
                    console.log("Game state set to: " + WaitingInLobby.STATE);
                    this.storeService.setState(state);
                } else if (gameState === "wordSelect") {
                    const playerDrawingUuid = payload.requestCurrentGameInfo.selectedClient.UUID;
                    const playerDrawing = this.storeService.getPlayerWithUuid(playerDrawingUuid)
                    console.log("Game state set to: " + WaitingForPlayerToChooseWord.STATE);
                    const remainingTime = payload.requestCurrentGameInfo.timerRemaining;
                    const state = new WaitingForPlayerToChooseWord(
                        playerDrawing,
                        this._convertNanoSecsToSecs(remainingTime)
                    );
                    this.storeService.setState(state);
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
        this.storeService.setPlayers(players)
    }

    onGameEvent(data) {
        const payload = data.payload;
        if (payload.gameMasterAPI === "waitForStart") {
            this.storeService.setPlayerReadyState(payload.waitForStart);
        }
        else if (payload.gameMasterAPI === "wordSelect") {
            this.storeService.setRoundNumber(payload.wordSelect.round);
            const playerUuid = this.storeService.getPlayerUuid();
            if (playerUuid === payload.wordSelect.chosenUUID) {
                const player = this.storeService.getPlayerWithUuid(playerUuid);
                const state = new ChoosingWord(
                    player,
                    payload.wordSelect.choices,
                    this._convertNanoSecsToSecs(payload.wordSelect.duration)
                );
                console.log("Setting game state to ChoosingWord")
                this.storeService.setState(state);
                EventBus.$emit(Event.TIMER_RESET, state.duration);
            } else {
                const player = this.storeService.getPlayerWithUuid(playerUuid)
                const state = new WaitingForPlayerToChooseWord(
                    player,
                    this._convertNanoSecsToSecs(payload.wordSelect.duration)
                );
                console.log("Setting game state to WaitingForPlayerToChooseWord")
                this.storeService.setState(state);
                EventBus.$emit(Event.TIMER_RESET, state.duration);
            }
        } else if (payload.gameMasterAPI === "playTime") {
            const currentState = this.storeService.getState();
            if (currentState.state === ChoosingWord.STATE) {
                const selectedWord = this.storeService.getWordSelected();
                const state = new Drawing(
                    selectedWord,
                    this._convertNanoSecsToSecs(payload.playTimeSend.duration)
                );
                console.log("Setting game state to Drawing")
                this.storeService.setState(state);
                EventBus.$emit(Event.TIMER_RESET, state.duration);
            } else if (currentState.state === WaitingForPlayerToChooseWord.STATE) {
                const state = new Guessing(
                    payload.playTimeSend.hint,
                    this._convertNanoSecsToSecs(payload.playTimeSend.duration)
                );
                console.log("Setting game state to Guessing")
                this.storeService.setState(state);
                EventBus.$emit(Event.TIMER_RESET, state.duration);
            }
        } else if (payload.gameMasterAPI === "playTime") {
            const currentState = this.storeService.getState();
            if (currentState.state === ChoosingWord.STATE) {
                const selectedWord = this.storeService.getWordSelected();
                const state = new Drawing(
                    selectedWord,
                    this._convertNanoSecsToSecs(payload.playTimeSend.duration)
                );
                console.log("Setting game state to Drawing")
                this.storeService.setState(state);
            } else if (currentState.state === WaitingForPlayerToChooseWord.STATE) {
                const state = new Guessing(
                    payload.playTimeSend.hint,
                    this._convertNanoSecsToSecs(payload.playTimeSend.duration)
                );
                console.log("Setting game state to Guessing")
                this.storeService.setState(state);
            }
        }
        // else if (payload.gameMasterAPI === "playTime") {
        //     const currentState = this.storeService.getState();
        //     if (currentState.state === ChoosingWord.STATE) {
        //         const selectedWord = this.storeService.getWordSelected();
        //         const state = new Drawing(
        //             selectedWord,
        //             this._convertNanoSecsToSecs(payload.playTimeSend.duration)
        //         );
        //         this.storeService.setState(state);
        //     }
        //     else if (currentState.state === WaitingForPlayerToChooseWord.STATE) {
        //         const state = new Guessing(
        //             payload.playTimeSend.hint,
        //             this._convertNanoSecsToSecs(payload.playTimeSend.duration)
        //         );
        //         this.storeService.setState(state);
        //     } else {
        //         const totalScores = payload.playTimeSend.totalScore;
        //         Object.keys(totalScores).forEach((key) => {
        //             const playerScore = {
        //                 uuid: key,
        //                 score: totalScores[key],
        //             };
        //             this.storeService.setPlayerScore(playerScore);
        //         });

        //         let correctClientUuid = payload.playTimeSend.correctClient.UUID;

        //         const player = this.storeService.getPlayerWithUuid(correctClientUuid);
        //         EventBus.$emit(Event.CORRECT_GUESS, player);
        //     }
        // }
        // else if (payload.gameMasterAPI === "scoreTime") {
        //     this.storeService.setRoundNumber(payload.scoreTime.round);
        // } else if (payload.gameMasterAPI === "showResults") {
        //     const state = new GameOver();
        //     this.storeService.setState(state);
        // }

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
                EventBus.$emit(Event.PLAYER_JOIN, player);
            });

            const playersLeft = this.getPlayerListDifference(
                currentPlayers,
                players
            );
            playersLeft.forEach((player) => {
                console.log("Player left");
                EventBus.$emit(Event.PLAYER_LEFT, player);
            });
        }
    }

    getPlayerListDifference(playerList1, playerList2) {
        return playerList1.filter(
            (p) => !playerList2.map((q) => q.uuid).includes(p.uuid)
        );
    }

    // handleEvent(data) {
    //     console.log("handling event");
    //     const api = data.api;
    //     const payload = data.payload;

    //     if (api === "hub") {
    //         const players = this.mapToPlayers(payload);
    //         this.handlePlayersChanged(players);
    //         this.setPlayers(players);
    //     } else if (api === "chat") {
    //         if (payload.history) {
    //             EventBus.$emit(Event.CHAT_HISTORY, payload.history);
    //         }
    //         else if (payload.message) {
    //             EventBus.$emit(Event.CHAT_MESSAGE, payload.message);
    //         }
    //     } else if (api === "game") {
    //         if (payload.gameMasterAPI === "waitForStart") {
    //             this.storeService.setPlayerReadyState(payload.waitForStart);
    //         }
    //         else if (payload.gameMasterAPI === "wordSelect") {
    //             this.storeService.setRoundNumber(payload.wordSelect.round);
    //             const playerUuid = this.playerUuid;
    //             if (playerUuid === payload.wordSelect.chosenUUID) {
    //                 const player = this.storeService.getPlayerWithUuid(playerUuid);
    //                 const state = new ChoosingWord(
    //                     player,
    //                     payload.wordSelect.choices,
    //                     this._convertNanoSecsToSecs(payload.wordSelect.duration)
    //                 );
    //                 console.log("Setting game state to ChoosingWord")
    //                 this.storeService.setState(state);
    //             } else {
    //                 const player = this.storeService.getPlayerWithUuid(playerUuid)
    //                 const state = new WaitingForPlayerToChooseWord(
    //                     player,
    //                     this._convertNanoSecsToSecs(payload.wordSelect.duration)
    //                 );
    //                 console.log("Setting game state to WaitingForPlayerToChooseWord")
    //                 this.storeService.setState(state);
    //             }
    //         }
    //         else if (payload.gameMasterAPI === "playTime") {
    //             const currentState = this.storeService.getState();
    //             if (currentState.state === ChoosingWord.STATE) {
    //                 const selectedWord = this.storeService.getWordSelected();
    //                 const state = new Drawing(
    //                     selectedWord,
    //                     this._convertNanoSecsToSecs(payload.playTimeSend.duration)
    //                 );
    //                 this.storeService.setState(state);
    //             }
    //             else if (currentState.state === WaitingForPlayerToChooseWord.STATE) {
    //                 const state = new Guessing(
    //                     payload.playTimeSend.hint,
    //                     this._convertNanoSecsToSecs(payload.playTimeSend.duration)
    //                 );
    //                 this.storeService.setState(state);
    //             } else {
    //                 const totalScores = payload.playTimeSend.totalScore;
    //                 Object.keys(totalScores).forEach((key) => {
    //                     const playerScore = {
    //                         uuid: key,
    //                         score: totalScores[key],
    //                     };
    //                     this.storeService.setPlayerScore(playerScore);
    //                 });

    //                 let correctClientUuid = payload.playTimeSend.correctClient.UUID;

    //                 const player = this.storeService.getPlayerWithUuid(correctClientUuid);
    //                 EventBus.$emit(Event.CORRECT_GUESS, player);
    //             }
    //         }
    //         else if (payload.gameMasterAPI === "scoreTime") {
    //             this.storeService.setRoundNumber(payload.scoreTime.round);
    //         } else if (payload.gameMasterAPI === "showResults") {
    //             const state = new GameOver();
    //             this.storeService.setState(state);
    //         }

    //     }
    // }

    _convertNanoSecsToSecs(durationNS) {
        return Math.round(durationNS / NANOSECOND_TO_SECONDS_FACTOR);
    }
}

