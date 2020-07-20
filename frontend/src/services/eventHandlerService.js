import { EventBus } from "../eventBus.js";
import EventHandlerFactory from "./eventHandlers/eventHandlerFactory";

// Service for handling websocket events
export default class EventHandlerService {
    handle(api, payload) {
        const eventHandlerFactory = new EventHandlerFactory();
        const eventHandler = eventHandlerFactory.getHandler(api);
        if (eventHandler != null) {
            console.log("----- Event Received -----");
            console.log(api);
            console.log(payload);
            eventHandler.handle(payload);
        }

        EventBus.$emit(api, payload);
    }
}
