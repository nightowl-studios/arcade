export default class ApplicationController {
    constructor(apiServiceFacade, storeService) {
        this.apiServiceFacade = apiServiceFacade;
        this.storeService = storeService;
        this.webSocketService = apiServiceFacade.getWebSocketService();
    }

    async initApplication(lobbyId) {
        if (this.storeService.getLobbyId() === lobbyId) {
            this.webSocketService.createConnection(lobbyId);
            return true;
        } else {
            if (await this.checkLobbyExists(lobbyId)) {
                this.storeService.setLobbyId(lobbyId);
                this.webSocketService.createConnection(lobbyId);
                return true;
            } else {
                return false;
            }
        }
    }

    async createLobby() {
        return this.apiServiceFacade.createLobby();
    }

    async checkLobbyExists(lobbyId) {
        return this.apiServiceFacade.checkLobbyExists(lobbyId);
    }

    setLobbyId(lobbyId) {
        this.storeService.setLobbyId(lobbyId);
    }

    changeNickname(nickname) {
        if (this.webSocketService.getConnection() != null &&
            this.webSocketService.getConnection().isConnected()) {
            this.apiServiceFacade.changeNickname(nickname);
        }

        this.storeService.setNickname(nickname);
    }
}
