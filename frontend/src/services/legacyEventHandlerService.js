import { EventBus } from "../eventBus.js";
import EventHandlerFactory from "./eventHandlers/eventHandlerFactory";

// Service for handling websocket events
export default class EventHandlerService {
    constructor() {
        this.eventHandlerFactory = new EventHandlerFactory();
    }

    handle(api, payload) {
        const eventHandler = this.eventHandlerFactory.getHandler(api);
        if (eventHandler != null) {
            console.log("----- Event Received -----");
            console.log(api);
            console.log(payload);
            eventHandler.handle(payload);
        }

        EventBus.$emit(api, payload);
    }
}
