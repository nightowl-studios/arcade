import { EventBus } from "@/eventBus";
import { Event } from "@/events";
import { WebSocketEvent } from "../../communication/webSocketEvent";
import GameManager from "../../scribble/gameManager";

export default class ApiReceiverService {
    constructor(applicationController, gameController, applicationStoreService, gameStoreService) {
        this.applicationControler = applicationController;
        this.gameController = gameController
        this.applicationStoreService = applicationStoreService;
        this.gameStoreService = gameStoreService;
    }

    update(event, data) {
        console.log("Received an event")
        console.log('Event:', event);
        console.log('Data:', data);

        if (event === WebSocketEvent.WEBSOCKET_CONNECTED) {
            EventBus.$emit(Event.WEBSOCKET_CONNECTED, data);
            this.gameController.initGame();
            this.applicationControler.changeNickname(this.applicationStoreService.getNickname());
        }

        if (event === WebSocketEvent.WEBSOCKET_ONMESSAGE) {
            if (data.api === "auth") {
                this.applicationStoreService.setPlayerUuid(data.payload.uuid)
            }

            const gameManager = new GameManager(this.gameStoreService);
            const currentState = gameManager.getCurrentState();
            if (currentState == null) {
                gameManager.initGame(data);

                if (this.gameStoreService.getState() != null) {
                    this.gameStoreService.setLoading(false);
                }
            } else {
                gameManager.handleEvent(data);
            }
        }
    }
}
