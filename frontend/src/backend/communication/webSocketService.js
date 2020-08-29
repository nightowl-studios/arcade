import WebSocketConnection from "./webSocketConnection";

export default class WebSocketService {
    constructor(webSocketUrl) {
        this.webSocketUrl = webSocketUrl;
        this.webSocketConnection = new WebSocketConnection();
    }

    createConnection(lobbyId) {
        if (!this.webSocketConnection.isConnected()) {
            this.webSocketConnection.connect(this.webSocketUrl, lobbyId);
        } else {
            console.error("There is an existing connection already");
        }
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
