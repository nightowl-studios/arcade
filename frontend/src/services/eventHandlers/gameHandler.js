import { EventBus } from "@/eventBus.js";
import { Event } from "@/events";
import ChoosingWord from "@/modules/scribble/stores/states/gamestates";
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
            }
            // TODO store payload in vuex
            store.commit("scribble/setChosenUuid", payload.wordSelect.chosenUUID);
            store.commit("scribble/setIsCanvasLocked", payload.wordSelect.lockCanvas);
            EventBus.$emit(Event.START_GAME);
        }
    }
}
