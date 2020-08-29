import { store } from "@/store";
export default class ApplicationStoreService {
    setLobbyId(lobbyId) {
        store.commit("application/setLobbyId", lobbyId);
    }

    getLobbyId() {
        return store.getters["application/getLobbyId"];
    }

    setPlayerUuid(uuid) {
        store.commit("application/setPlayerUuid", uuid);
    }

    getPlayerUuid() {
        return store.getters["application/getPlayerUuid"];
    }

    setNickname(nickname) {
        store.commit("application/setNickname", nickname);
    }

    getNickname() {
        return store.getters["application/getNickname"];
    }
}
