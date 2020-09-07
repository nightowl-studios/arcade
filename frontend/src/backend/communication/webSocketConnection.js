import { WebSocketEvent } from "./webSocketEvent";

export default class WebSocketConnection {
    constructor() {
        this.listeners = [];
        this.webSocket = null;
    }

    connect(url, lobbyId) {
        const webSocketUrl = `${url}/${lobbyId}`;
        this.url = webSocketUrl;
        this.lobbyId = lobbyId;
        if (this.webSocket == null) {
            this.webSocket = new WebSocket(webSocketUrl);
            this.init();
        } else {
            console.log("There is an existing websocket");
        }
    }

    disconnect() {
        this.webSocket.close();
        this.webSocket = null;
    }

    init() {
        this.webSocket.onopen = () => {
            console.log(
                `Successfully connected to the websocket. url: ${this.url}`
            );

            // TODO add this legacy code somehow
            // const arcadeSession = this.cookieService.getArcadeCookie();
            // if (arcadeSession != null && arcadeSession.ContainsToken != false) {
            //     this.send(arcadeSession);
            // } else {
            //     const noToken = {
            //         api: "auth",
            //         payload: {
            //             ContainsToken: false,
            //         },
            //     };
            //     this.send(noToken);
            // }

            let eventType = WebSocketEvent.WEBSOCKET_CONNECTED;
            let data = this.lobbyId;
            this.listeners.forEach(listener => listener.update(eventType, data))
        }

        this.webSocket.onmessage = (event) => {
            let eventType = WebSocketEvent.WEBSOCKET_ONMESSAGE;
            let data = JSON.parse(event.data);

            this.listeners.forEach(listener => listener.update(eventType, data))
        }

        this.webSocket.onclose = () => {
            let eventType = WebSocketEvent.WEBSOCKET_CLOSED;
            let data = JSON.parse(event.data);

            this.listeners.forEach(listener => listener.update(eventType, data))
        };
    }

    send(data) {
        if (this.isConnected()) {
            const json = JSON.stringify(data);
            this.webSocket.send(json);
        } else {
            console.error("Unable to send data. Not Connected");
        }
    }

    isConnected() {
        return (
            this.webSocket != null &&
            this.webSocket.readyState === WebSocket.OPEN
        );
    }

    addListener(listener) {
        this.listeners.push(listener);
    }
}
