import { EventBus } from "@/eventBus.js";
import { Event } from "@/events";
import { ChoosingWord, WaitingForPlayerToChooseWord } from "@/modules/scribble/stores/states/gamestates";
import { store } from "@/store";

// Event handler for Game API
export default class GameHandler {
    handle(payload) {
        if (payload.gameMasterAPI === "wordSelect") {
            const playerUuid = store.getters["application/getPlayerUuid"]
            if (playerUuid === payload.wordSelect.chosenUUID) {
                const player = store.getters["application/getPlayerWithUuid"](playerUuid);
                const state = new ChoosingWord(player, payload.wordSelect.choices);
                store.commit("scribble/setGameState", state);
            } else {
                const player = store.getters["application/getPlayerWithUuid"](playerUuid);
                const state = new WaitingForPlayerToChooseWord(player);
                store.commit("scribble/setGameState", state);
            }

            EventBus.$emit(Event.START_GAME);
        }
    }
}
