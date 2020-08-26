import { createChatMessage } from "@/utility/WebSocketMessageUtils";
// Service for making request to chat API.
export default class GameApiService {
    sendChatMessage(webSocketConnection, message) {
        const data = createChatMessage(message);
        this.sendMessage(webSocketConnection, data)
    }

    requestChatHistory(webSocketConnection) {
        const request = {
            api: "chat",
            payload: {
                requestHistory: true,
            },
        };
        this.sendMessage(webSocketConnection, request);
    }

    sendMessage(webSocketConnection, data) {
        webSocketConnection.send(data);
    }
}
