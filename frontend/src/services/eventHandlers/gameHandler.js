import { EventBus } from "@/eventBus";
import {
    ChoosingWord,
    Drawing,
    Guessing,
    WaitingForPlayerToChooseWord,
} from "@/modules/scribble/stores/states/gamestates";
import { store } from "@/store";

// Event handler for Game API
export default class GameHandler {
    handle(payload) {
        if (payload.gameMasterAPI === "wordSelect") {
            const playerUuid = store.getters["application/getPlayerUuid"];
            if (playerUuid === payload.wordSelect.chosenUUID) {
                const player = store.getters["application/getPlayerWithUuid"](
                    playerUuid
                );
                const state = new ChoosingWord(
                    player,
                    payload.wordSelect.choices
                );
                store.commit("scribble/setGameState", state);
            } else {
                const player = store.getters["application/getPlayerWithUuid"](
                    playerUuid
                );
                const state = new WaitingForPlayerToChooseWord(player);
                store.commit("scribble/setGameState", state);
            }

            EventBus.$emit(Event.START_GAME);
        } else if (payload.gameMasterAPI === "playTime") {
            const currentState = store.getters["scribble/getCurrentState"];
            if (currentState.state === ChoosingWord.STATE) {
                const state = new Drawing();
                store.commit("scribble/setGameState", state);
            } else if (
                currentState.state === WaitingForPlayerToChooseWord.STATE
            ) {
                const state = new Guessing();
                store.commit("scribble/setGameState", state);
            }
        }
    }
}
