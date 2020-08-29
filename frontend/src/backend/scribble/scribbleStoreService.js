import { store } from "@/store";
export default class ScribbleStoreService {
    getState() {
        return store.getters["scribble/getGameState"];
    }

    setState(state) {
        store.commit("scribble/setGameState", state);
    }

    setPlayers(players) {
        store.commit("scribble/setPlayers", players);
    }

    setLoading(isLoading) {
        store.commit("scribble/setLoading", isLoading);
    }
}
