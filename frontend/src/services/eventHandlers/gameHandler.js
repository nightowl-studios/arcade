import { EventBus } from "@/eventBus.js";
import { Event } from "@/events";

// Event handler for Game API
export default class GameHandler {
    constructor() {}

    handle(payload) {
        if (payload.gameMasterAPI === "playerSelect") {
            EventBus.$emit(Event.START_GAME);
        }
    }
}
