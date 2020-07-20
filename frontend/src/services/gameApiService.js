// Service for making request to game API.
export default class GameApiService {
    constructor(webSocketService) {
        this.webSocketService = webSocketService;
    }

    startGame() {
        const message = {
            api: "game",
            payload: {
                gameMasterAPI: "waitForStart",
                waitForStart: {
                    startGame: true,
                },
            },
        };
        this.webSocketService.send(message);
    }
}
