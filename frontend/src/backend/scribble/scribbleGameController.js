
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

    selectWord(index, word) {
        this.storeService.setWordSelected(word);
        this.apiServiceFacade.selectWord(index);
    }

    sendChatMessage(msg) {
        this.apiServiceFacade.sendChatMessage(msg);
    }

    draw(drawAction) {
        this.apiServiceFacade.draw(drawAction);
    }

    changeNickname(nickname) {
        this.apiServiceFacade.changeNickname(nickname);
    }

    authenticate() {
        this.apiServiceFacade.authenticate();
    }

    resetStoreState() {
        this.storeService.resetState();
    }

}
