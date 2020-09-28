import { store } from "@/store";
export default class ApplicationStoreService {
    setLobbyId(lobbyId) {
        store.commit("application/setLobbyId", lobbyId);
    }

    getLobbyId() {
        return store.getters["application/getLobbyId"];
    }

    setNickname(nickname) {
        store.commit("application/setNickname", nickname);
    }

    getNickname() {
        return store.getters["application/getNickname"];
    }

    resetState() {
        store.commit("application/resetState");
    }
}
