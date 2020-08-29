class ApiSenderFacade {
    setWebSocketService(webSocketService) {
        this.webSocketService = webSocketService;
    }

    setHubApiService(hubApiService) {
        this.hubApiService = hubApiService;
    }

    setGameApiService(gameApiService) {
        this.gameApiService = gameApiService;
    }

    createLobby() {
        return this.hubApiService.createLobby();
    }

    checkLobbyExists(lobbyId) {
        return this.hubApiService.checkLobbyExists(lobbyId);
    }

    setIsReady(isReady) {
        this.gameApiService.setIsReady(this.webSocketService.getConnection(), isReady)
    }

    // selectWordToDraw(index) {

    // }

    requestCurrentGameInfo() {
        this.gameApiService.requestCurrentGameInfo(this.webSocketService.getConnection());
    }

    // sendChatMessage(message) {

    // }

    // requestChatHistory() {

    // }

    changeNickname(nickname) {
        this.hubApiService.changeNickname(this.webSocketService.getConnection(), nickname);
    }

    getWebSocketService() {
        return this.webSocketService;
    }
}

const singletonInstance = new ApiSenderFacade;

export default singletonInstance;

