export default class DrawApiService {
    constructor(webSocketService) {
        this.webSocketService = webSocketService;
    }

    requestDrawHistory() {
        const request = {
            api: "draw",
            payload: {
                requestHistory: true
            }
        }

        this.sendMessage(request);
    }

    draw(drawAction) {
        const request = {
            api: "draw",
            payload: {
                action: drawAction,
                requestHistory: false,
            }
        }

        this.sendMessage(request);
    }

    sendMessage(data) {
        this.webSocketService.send(data);
    }
}
