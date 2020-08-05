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
        this.webSocketService.send(message);
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
        this.webSocketService.send(message);
    }
}
