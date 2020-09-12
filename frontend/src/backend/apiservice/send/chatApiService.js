// Service for making request to chat API.
export default class GameApiService {
    constructor(webSocketService) {
        this.webSocketSerivce = webSocketService;
    }

    sendChatMessage(message) {
        const data = {
            api: "chat",
            payload: {
                message: message
            }
        }

        this.sendMessage(data)
    }

    requestChatHistory() {
        const request = {
            api: "chat",
            payload: {
                requestHistory: true,
            },
        };
        this.sendMessage(request);
    }

    sendMessage(data) {
        this.webSocketSerivce.send(data);
    }
}
