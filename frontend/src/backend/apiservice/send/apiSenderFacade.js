export default class ApiSenderFacade {
    constructor(webSocketService, hubApiService) {
        this.webSocketService = webSocketService;
        this.hubApiService = hubApiService;
        // this.gameApiService;
        // this.chatApiService;
        // this.hubApiService;
    }

    createLobby() {
        return this.hubApiService.createLobby();
    }

    // checkLobbyExists(lobbyId) {

    // }

    // setPlayerIsReadyState(isReady) {

    // }

    // selectWordToDraw(index) {

    // }

    // requestCurrentGameInfo() {

    // }

    // sendChatMessage(message) {

    // }

    // requestChatHistory() {

    // }
}
