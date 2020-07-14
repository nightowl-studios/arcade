import { EventBus } from '../eventBus.js';

export default class WebSocketService {
    constructor(webSocketURL) {
        this.webSocketURL = webSocketURL;
        this.webSocket = null;
    }

    getWebSocketURL() {
        return this.webSocketURL;
    }

    connect(lobbyId) {
        console.log("Connecting to websocket...");
        let webSocketURL = this.webSocketURL + "/" + lobbyId;
        this.webSocket = new WebSocket(webSocketURL)
        this.initWebSocket(this.webSocket, lobbyId);
    }

    initWebSocket(webSocket, lobbyId) {
        webSocket.onopen = () => {
            console.log("Successfully connected to the websocket. ID: " + lobbyId);
            EventBus.$emit('connected', lobbyId);
        }

        webSocket.onmessage = (event) => {
            let json = JSON.parse(event.data);
            let apiName = json.api;
            EventBus.$emit(apiName, json.payload);
        }
    }

    send(data) {
        if (this.isConnected()) {
            let json = JSON.stringify(data);
            this.webSocket.send(json);
        } else {
            console.error("NOT CONNECTED");
        }
    }

    isConnected() {
        return this.webSocket != null && this.webSocket.readyState === WebSocket.OPEN;
    }
}