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
                let json = JSON.parse(event.data);
                let apiName = json.api;

                this.$emit(apiName, json.payload);
            }
        }
    }
});