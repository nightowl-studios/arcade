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

    setAuthApiService(authApiService) {
        this.authApiService = authApiService;
    }

    createLobby() {
        return this.hubApiService.createLobby();
    }

    checkLobbyExists(lobbyId) {
        return this.hubApiService.checkLobbyExists(lobbyId);
    }

    setIsReady(isReady) {
        this.gameApiService.setIsReady(isReady)
    }

    // selectWordToDraw(index) {

    // }

    requestCurrentGameInfo() {
        this.gameApiService.requestCurrentGameInfo();
    }

    sendChatMessage(msg) {
        this.chatApiService.sendChatMessage(msg)
    }

    requestChatHistory() {
        this.chatApiService.requestChatHistory();
    }

    changeNickname(nickname) {
        this.hubApiService.changeNickname(nickname);
    }

    getWebSocketService() {
        return this.webSocketService;
    }

    requestDrawHistory() {
        this.drawApiService.requestDrawHistory();
    }

    selectWord(index) {
        this.gameApiService.selectWord(index);
    }

    draw(drawAction) {
        this.drawApiService.draw(drawAction);
    }

    authenticate() {
        this.authApiService.authenticate();
    }
}

const singletonInstance = new ApiSenderFacade;

export default singletonInstance;

