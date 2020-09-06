import Player from "./entities/player";
import {
    ChoosingWord,
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

        if (api === "hub") {
            console.log("loading players...");
            this.setPlayers(payload);

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
            this.setPlayers(payload);
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
        }
    }

    _convertNanoSecsToSecs(durationNS) {
        return durationNS / NANOSECOND_TO_SECONDS_FACTOR;
    }

    setPlayers(payload) {
        const players = payload.connectedClients.map(
            (client) =>
                new Player(
                    client.clientUUID.UUID,
                    client.nickname,
                    client.joinOrder
                )
        );

        this.storeService.setPlayers(players);
    }
}

