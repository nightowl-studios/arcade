import { EventBus } from "@/eventBus.js";
import { Event } from "@/events";
import { store } from "@/store";

// Event handler for Game API
export default class GameHandler {
    handle(payload) {
        if (payload.gameMasterAPI === "playerSelect") {
            // TODO store payload in vuex
            store.commit("scribble/setChosenUuid", payload.playerSelect.chosenUUID);
            EventBus.$emit(Event.START_GAME);
        }
    }
}
