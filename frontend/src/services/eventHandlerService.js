import { EventBus } from "../eventBus.js";
import EventHandlerFactory from "./eventHandlers/eventHandlerFactory";

// Service for handling websocket events
export default class EventHandlerService {
    constructor() {}

    handle(api, payload) {
        console.log("----- Event Received -----");
        console.log(api);
        console.log(payload);

        let eventHandlerFactory = new EventHandlerFactory();
        let eventHandler = eventHandlerFactory.getHandler(api);
        if (eventHandler != null) {
            eventHandler.handle(payload);
        }

        EventBus.$emit(api, payload);
    }
}
