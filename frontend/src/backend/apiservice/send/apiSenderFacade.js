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

    setChatApiService(chatApiService) {
        this.chatApiService = chatApiService;
    }

    setDrawApiService(drawApiService) {
        this.drawApiService = drawApiService;
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
        console.log("requesting current game info");
        this.gameApiService.requestCurrentGameInfo(this.webSocketService.getConnection());
    }

    // sendChatMessage(message) {

    // }

    requestChatHistory() {
        this.chatApiService.requestChatHistory(this.webSocketService.getConnection());
    }

    changeNickname(nickname) {
        this.hubApiService.changeNickname(this.webSocketService.getConnection(), nickname);
    }

    getWebSocketService() {
        return this.webSocketService;
    }

    requestDrawHistory() {
        this.drawApiService.requestDrawHistory(this.webSocketService.getConnection());
    }
}

const singletonInstance = new ApiSenderFacade;

export default singletonInstance;

