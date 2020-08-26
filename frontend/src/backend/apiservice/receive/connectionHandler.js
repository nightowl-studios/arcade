import ApiSenderFacade from "@/backend/apiservice/send/apiSenderFacade";

export default class ConnectionHandler {
    onNewConnection() {
        ApiSenderFacade.requestCurrentGameInfo();
        // request game details
        // Populate vuex with game etails
    }
}
