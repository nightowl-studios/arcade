import { store } from "@/store";
export default class StoreService {
    getState() {
        return store.getters["scribble/getGameState"];
    }

    setState(state) {
        store.commit("scribble/setGameState", state);
    }

    setPlayers(players) {
        store.commit("application/setPlayers", players);
    }

    setPlayerUuid(uuid) {
        store.commit("application/setPlayerUuid", uuid);
    }
}
