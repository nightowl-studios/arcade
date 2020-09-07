import { EventBus } from "@/eventBus";
import { Event } from "@/events";
import Player from "./entities/player";
import {
    ChoosingWord,
    Drawing,

    GameOver, Guessing,
    WaitingForPlayerToChooseWord, WaitingInLobby
} from "./states/gameStates";

const NANOSECOND_TO_SECONDS_FACTOR = 1000000000;

export default class GameManager {
    constructor(playerUuid, gameController, storeService) {
        this.playerUuid = playerUuid;
        this.gameController = gameController;
        this.storeService = storeService;
    }

    getCurrentState() {
        return this.storeService.getState();
    }

    initGame(data) {
        const api = data.api;
        const payload = data.payload;

        if (api === "auth") {
            // TODO
            // const cookieService = new CookieService();
            // cookieService.setArcadeCookie(payload.tokenMessage);
            // store.commit("application/setPlayerUuid", payload.uuid);
        }

        if (api === "hub") {
            console.log("loading players...");
            const players = this.mapToPlayers(payload);
            this.setPlayers(players);

            this.gameController.initGame();
        }

        if (api === "game") {
            if (payload.gameMasterAPI === "requestCurrentGameInfo") {
                console.log("Initializing games state...");
                const gameState = payload.requestCurrentGameInfo.gameState;
                if (gameState === "waitForStart") {
                    const state = new WaitingInLobby();
                    this.storeService.setState(state);
                } else if (gameState === "wordSelect") {
                    const state = new WaitingForPlayerToChooseWord();
                    this.storeService.setState(state);
                }
            }
        }
    }

    handleEvent(data) {
        console.log("handling event");
        const api = data.api;
        const payload = data.payload;

        if (api === "hub") {
            const players = this.mapToPlayers(payload);
            this.handlePlayersChanged(players);
            this.setPlayers(players);
        } else if (api === "chat") {
            if (payload.history) {
                EventBus.$emit(Event.CHAT_HISTORY, payload.history);
            }
            else if (payload.message) {
                EventBus.$emit(Event.CHAT_MESSAGE, payload.message);
            }
        } else if (api === "game") {
            if (payload.gameMasterAPI === "waitForStart") {
                this.storeService.setPlayerReadyState(payload.waitForStart);
            }
            else if (payload.gameMasterAPI === "wordSelect") {
                this.storeService.setRoundNumber(payload.wordSelect.round);
                const playerUuid = this.playerUuid;
                if (playerUuid === payload.wordSelect.chosenUUID) {
                    const player = this.storeService.getPlayerWithUuid(playerUuid);
                    const state = new ChoosingWord(
                        player,
                        payload.wordSelect.choices,
                        this._convertNanoSecsToSecs(payload.wordSelect.duration)
                    );
                    console.log("Setting game state to ChoosingWord")
                    this.storeService.setState(state);
                } else {
                    const player = this.storeService.getPlayerWithUuid(playerUuid)
                    const state = new WaitingForPlayerToChooseWord(
                        player,
                        this._convertNanoSecsToSecs(payload.wordSelect.duration)
                    );
                    console.log("Setting game state to WaitingForPlayerToChooseWord")
                    this.storeService.setState(state);
                }
            }
            else if (payload.gameMasterAPI === "playTime") {
                const currentState = this.storeService.getState();
                if (currentState.state === ChoosingWord.STATE) {
                    const selectedWord = this.storeService.getWordSelected();
                    const state = new Drawing(
                        selectedWord,
                        this._convertNanoSecsToSecs(payload.playTimeSend.duration)
                    );
                    this.storeService.setState(state);
                }
                else if (currentState.state === WaitingForPlayerToChooseWord.STATE) {
                    const state = new Guessing(
                        payload.playTimeSend.hint,
                        this._convertNanoSecsToSecs(payload.playTimeSend.duration)
                    );
                    this.storeService.setState(state);
                } else {
                    const totalScores = payload.playTimeSend.totalScore;
                    Object.keys(totalScores).forEach((key) => {
                        const playerScore = {
                            uuid: key,
                            score: totalScores[key],
                        };
                        this.storeService.setPlayerScore(playerScore);
                    });

                    let correctClientUuid = payload.playTimeSend.correctClient.UUID;

                    const player = this.storeService.getPlayerWithUuid(correctClientUuid);
                    EventBus.$emit(Event.CORRECT_GUESS, player);
                }
            }
            else if (payload.gameMasterAPI === "scoreTime") {
                this.storeService.setRoundNumber(payload.scoreTime.round);
            } else if (payload.gameMasterAPI === "showResults") {
                const state = new GameOver();
                this.storeService.setState(state);
            }

        }
    }

    _convertNanoSecsToSecs(durationNS) {
        return durationNS / NANOSECOND_TO_SECONDS_FACTOR;
    }

    setPlayers(players) {
        this.storeService.setPlayers(players);
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
}

