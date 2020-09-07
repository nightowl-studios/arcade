import { EventBus } from "@/eventBus.js";
import { Event } from "@/events";

// Event handler for Chat API
export default class ChatHandler {
    handle(payload) {
        if (payload.history) {
            EventBus.$emit(Event.CHAT_HISTORY, payload.history);
        }
        else if (payload.message) {
            EventBus.$emit(Event.CHAT_MESSAGE, payload.message);
        }
    }
}
