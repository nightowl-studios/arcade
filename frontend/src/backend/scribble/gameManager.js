import { WaitingInLobby } from "./states/gameStates";

export default class GameManager {
    constructor(storeService) {
        this.storeService = storeService;
    }

    getCurrentState() {
        return this.storeService.getState();
    }

    initGame(data) {
        console.log("init game");
        const api = data.api;
        const payload = data.payload;
        console.log(payload);

        if (api === "hub") {
            const state = new WaitingInLobby();
            this.storeService.setState(state);
        }
    }
}

