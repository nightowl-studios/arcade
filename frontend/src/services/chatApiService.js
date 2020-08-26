import { createChatMessage } from "../backend/apiservice/send/node_modules/@/utility/WebSocketMessageUtils";
// Service for making request to chat API.
export default class GameApiService {
    constructor(webSocketService) {
        this.webSocketService = webSocketService;
    }

    sendChatMessage(message) {
        const data = createChatMessage(message);
        this.webSocketService.send(data);
    }

    requestChatHistory() {
        const request = {
            api: "chat",
            payload: {
                requestHistory: true,
            },
        };
        this.webSocketService.send(request);
    }
}
