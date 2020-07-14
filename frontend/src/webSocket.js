import Vue from 'vue';
import VueCookies from 'vue-cookies'
import { EventBus } from './eventBus.js';

Vue.use(VueCookies);

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
                let arcadeSession = Vue.$cookies.get('arcade_session');
                if (arcadeSession != null &&
                    arcadeSession.ContainsToken != false) {
                    this.send(arcadeSession);
                } else {
                    let noToken = {
                        "api":"auth",
                        "payload":{
                            "ContainsToken": false
                        }
                    }
                    this.send(noToken);
                }

                console.log("Successfully connected to the websocket. ID: " + lobbyId);
                EventBus.$emit('connected', lobbyId);
            }

            webSocket.onmessage = (event) => {
                let json = JSON.parse(event.data);
                let apiName = json.api;

                if (apiName == "auth") {
                    Vue.$cookies.set("arcade_session", json.payload);
                }
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