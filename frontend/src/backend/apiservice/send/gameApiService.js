// Service for making request to game API.
export default class GameApiService {
    setIsReady(webSocketConnection, isReady) {
        const message = {
            api: "game",
            payload: {
                gameMasterAPI: "waitForStart",
                waitForStart: {
                    isReady: isReady,
                },
            },
        };

        this.sendMessage(webSocketConnection, message);
    }

    selectWord(webSocketConnection, index) {
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
        this.sendMessage(webSocketConnection, message);
    }

    requestCurrentGameInfo(webSocketConnection) {
        const message = {
            api: "game",
            payload: {
                gameMasterAPI: "requestCurrentGameInfo",
            },
        };
        this.sendMessage(webSocketConnection, message);
    }

    sendMessage(webSocketConnection, message) {
        webSocketConnection.send(message);
    }
}
