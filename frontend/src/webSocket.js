import Vue from 'vue';
import { EventBus } from './eventBus.js';

let webSocket = null;

function checkConnection() {
    return webSocket !== null;
}

export const ArcadeWebSocket = new Vue({
    methods: {
        connect: function(lobbyId) {
            console.log("Connecting to websocket...");
            let webSocketURL = this.$websocketURL + "/" + lobbyId;
            webSocket = new WebSocket(webSocketURL)

            webSocket.onopen = () => {
                console.log("Successfully connected to the websocket. ID: " + lobbyId);
                EventBus.$emit('connected', lobbyId);
            }

            webSocket.onmessage = (event) => {
                let json = JSON.parse(event.data);
                let apiName = json.api;
                EventBus.$emit(apiName, json.payload);
            }
        },
        send: function(data) {
            if (checkConnection()) {
                let json = JSON.stringify(data);
                webSocket.send(json);
            } else {
                console.error("NOT CONNECTED");
            }
        },
        isConnected: function() {
            return webSocket != null && webSocket.readyState === WebSocket.OPEN;
        }
    }
});