// Service for making request to game API.
export default class GameApiService {
    constructor(webSocketService) {
        this.webSocketService = webSocketService;
    }

    startGame() {
        let message = {
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
