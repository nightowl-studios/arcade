import { EventBus } from "@/eventBus";
import { Event } from "@/events";
import {
    ChoosingWord,
    Drawing,
    Guessing,
    WaitingForPlayerToChooseWord,
} from "@/modules/scribble/stores/states/gamestates";
import { store } from "@/store";

// Event handler for Game API
export default class GameHandler {
    constructor() {
        this.setGameStateKey = "scribble/setGameState";
    }

    handle(payload) {
        if (payload.gameMasterAPI === "wordSelect") {
            const playerUuid = store.getters["application/getPlayerUuid"];
            if (playerUuid === payload.wordSelect.chosenUUID) {
                const player = store.getters["application/getPlayerWithUuid"](
                    playerUuid
                );
                const state = new ChoosingWord(
                    player,
                    payload.wordSelect.choices,
                    payload.wordSelect.duration
                );
                store.commit(this.setGameStateKey, state);
            } else {
                const player = store.getters["application/getPlayerWithUuid"](
                    playerUuid
                );
                const state = new WaitingForPlayerToChooseWord(player);
                store.commit(this.setGameStateKey, state);
            }

            EventBus.$emit(Event.START_GAME);
        } else if (payload.gameMasterAPI === "playTime") {
            const currentState = store.getters["scribble/getGameState"];
            if (currentState.state === ChoosingWord.STATE) {
                const selectedWord = store.getters["scribble/getWordSelected"];
                const state = new Drawing(selectedWord);
                store.commit(this.setGameStateKey, state);
            } else if (
                currentState.state === WaitingForPlayerToChooseWord.STATE
            ) {
                const state = new Guessing(payload.playTimeSend.hint);
                store.commit(this.setGameStateKey, state);
            }
        }
    }
}
