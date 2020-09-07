// Service for making request to game API.
export default class GameApiService {
    constructor(webSocketService) {
        this.webSocketService = webSocketService;
    }

    setIsReady(isReady) {
        const message = {
            api: "game",
            payload: {
                gameMasterAPI: "waitForStart",
                waitForStart: {
                    isReady: isReady,
                },
            },
        };
        this.sendMessage(message);
    }

    selectWord(index) {
        const message = {
            api: "game",
            payload: {
                gameMasterAPI: "wordSelect",
                wordSelect: {
                    wordChosen: true,
                    choice: index,
                },
            },
        };
        this.sendMessage(message);
    }

    requestCurrentGameInfo() {
        const message = {
            api: "game",
            payload: {
                gameMasterAPI: "gameMasterAPI",
            },
        };
        this.sendMessage(message);
    }

    sendMessage(message) {
        this.webSocketService.send(message);
    }
}
