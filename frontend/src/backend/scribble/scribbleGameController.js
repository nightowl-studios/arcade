export default class ScribbleGameController {
    constructor(apiServiceFacade, storeService) {
        this.apiServiceFacade = apiServiceFacade;
        this.storeService = storeService;
    }

    initGame() {
        this.apiServiceFacade.requestCurrentGameInfo();
    }

    setIsReady(isReady) {
        this.apiServiceFacade.setIsReady(isReady)
    }

}
