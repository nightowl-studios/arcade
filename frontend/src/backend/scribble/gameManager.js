import Player from "./entities/player";
import { WaitingInLobby } from "./states/gameStates";

export default class GameManager {
    constructor(storeService) {
        this.storeService = storeService;
    }

    getCurrentState() {
        return this.storeService.getState();
    }

    initGame(data) {
        const api = data.api;
        const payload = data.payload;

        if (api === "hub") {
            console.log("initializing game");
            const state = new WaitingInLobby();
            this.storeService.setState(state);

            this.setPlayers(payload);
        }
    }

    handleEvent(data) {
        console.log("handling event");
        const api = data.api;
        const payload = data.payload;

        if (api === "hub") {
            this.setPlayers(payload);
        }
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

