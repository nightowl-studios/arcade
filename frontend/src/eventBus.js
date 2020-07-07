import Vue from 'vue';

let webSocket = null;

export const EventBus = new Vue({
    methods: {
        connect: function(lobbyId) {
            let webSocketURL = this.$websocketURL + "/" + lobbyId;
            console.log(webSocketURL);
            webSocket = new WebSocket(webSocketURL)

            webSocket.onopen = (event) => {
                console.log(event);
                console.log("Successfully connected to the websocket...");
                console.log(this);
                this.connectionState = "CONNECTED";
                this.$emit('connected');
            }

            webSocket.onmessage = (event) => {
                console.log("connection onMessage Triggered");
                let parseMsg = JSON.parse(event.data);
                let clients = parseMsg.payload.connectedClients;
                this.$emit('clientConnected', clients);
            }
        }
    }
});