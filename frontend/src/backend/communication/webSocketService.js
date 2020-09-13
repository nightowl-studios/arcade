
export default class WebSocketService {
    constructor(webSocketUrl, webSocketConnection) {
        this.webSocketUrl = webSocketUrl;
        this.webSocketConnection = webSocketConnection;
    }

    createConnection(lobbyId) {
        if (!this.webSocketConnection.isConnected()) {
            const url = `${this.webSocketUrl}/${lobbyId}`;
            this.webSocketConnection.connect(url);
        } else {
            console.error("There is an existing connection already");
        }
    }

    send(data) {
        this.webSocketConnection.send(data);
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
