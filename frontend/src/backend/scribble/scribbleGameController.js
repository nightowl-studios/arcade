export default class ScribbleGameController {
    constructor(apiServiceFacade, storeService) {
        this.apiServiceFacade = apiServiceFacade;
        this.storeService = storeService;
    }

    initGame() {
        console.log("requesting for current game info")
        this.apiServiceFacade.requestCurrentGameInfo();
    }

    setIsReady(isReady) {
        this.apiServiceFacade.setIsReady(isReady)
    }

    requestChatHistory() {
        this.apiServiceFacade.requestChatHistory();
    }

    requestDrawHistory() {
        this.apiServiceFacade.requestDrawHistory();
    }

}
