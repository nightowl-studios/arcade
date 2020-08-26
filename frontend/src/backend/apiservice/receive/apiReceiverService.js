import { EventBus } from "@/eventBus";
import { Event } from "@/events";
import { WebSocketEvent } from "../../communication/webSocketEvent";
import GameManager from "../../scribble/gameManager";
import StoreService from "../../scribble/storeService";
import ConnectionHandler from "./connectionHandler";

export default class ApiReceiverService {
    update(event, data) {
        console.log("Received an event")
        console.log('Event:', event);
        console.log('Data:', data);

        if (event === WebSocketEvent.WEBSOCKET_CONNECTED) {
            const connectionHandler = new ConnectionHandler();
            connectionHandler.onNewConnection();
            EventBus.$emit(Event.WEBSOCKET_CONNECTED, data);
        }

        if (event === WebSocketEvent.WEBSOCKET_ONMESSAGE) {
            const storeService = new StoreService();
            const gameManager = new GameManager(storeService);
            const currentState = gameManager.getCurrentState();
            if (currentState == null) {
                gameManager.initGame(data);
            } else {
                console.log("Handling event")
            }
        }
    }
}
