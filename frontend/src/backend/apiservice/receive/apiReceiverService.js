import { WebSocketEvent } from "../../communication/webSocketEvent";
import ConnectionHandler from "./connectionHandler";

export default class ApiReceiverService {
    update(event, data) {
        console.log(event);
        console.log(data);

        if (event === WebSocketEvent.WEBSOCKET_CONNECTED) {
            const connectionHandler = new ConnectionHandler();
            connectionHandler.onNewConnection();
        }
    }
}
