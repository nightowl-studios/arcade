import ApiReceiverService from "../apiservice/receive/apiReceiverService";
import WebSocketConnection from "./webSocketConnection";

export default class WebSocketService {
    constructor(webSocketUrl) {
        this.webSocketUrl = webSocketUrl;
        this.webSocketConnection = null;
    }

    createConnection(lobbyId) {
        if (this.webSocketConnection == null) {
            this.webSocketConnection = new WebSocketConnection();
            this.webSocketConnection.connect(this.webSocketUrl, lobbyId);

            this.init();
        } else {
            console.error("There is an existing connection already");
        }
    }

    init() {
        const apiReceiverService = new ApiReceiverService();
        this.webSocketConnection.addListener(apiReceiverService);
    }

    disconnect() {
        if (this.webSocketConnection != null && this.webSocketConnection.isConnected()) {
            this.webSocketConnection.disconnect();
        }
    }

    getConnection() {
        return this.webSocketConnection;
    }
}
