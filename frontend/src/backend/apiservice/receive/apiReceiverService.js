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
            const nickname = this.applicationStoreService.getNickname() == "" ? "RANDOM" : this.applicationStoreService.getNickname()
            this.applicationControler.changeNickname(nickname);
        }

        if (event === WebSocketEvent.WEBSOCKET_ONMESSAGE) {
            if (data.api === "auth") {
                this.applicationStoreService.setPlayerUuid(data.payload.uuid)
            }

            const playerUuid = this.applicationStoreService.getPlayerUuid();
            const gameManager = new GameManager(playerUuid, this.gameController, this.gameStoreService);
            const currentState = gameManager.getCurrentState();
            if (currentState == null) {
                this.initApplication(gameManager, data);
            } else {
                gameManager.handleEvent(data);
            }
        }
    }

    initApplication(gameManager, data) {
        gameManager.initGame(data);

        if (this.gameStoreService.getState() != null) {
            const nickname = this.applicationStoreService.getNickname() == "" ? "RANDOM" : this.applicationStoreService.getNickname()
            this.applicationControler.changeNickname(nickname);
            this.gameStoreService.setLoading(false);
        }
    }
}
