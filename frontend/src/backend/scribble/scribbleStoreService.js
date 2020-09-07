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

    setPlayerReadyState(readyState) {
        store.commit("scribble/setPlayerReadyState", readyState);
    }

    setRoundNumber(roundNumber) {
        store.commit("scribble/setRoundNumber", roundNumber);
    }

    getPlayers() {
        return store.getters["scribble/getPlayers"];
    }

    getPlayerWithUuid(uuid) {
        return store.getters["scribble/getPlayerWithUuid"](uuid);
    }

    setWordSelected(word) {
        store.commit("scribble/setWordSelected", word);
    }

    getWordSelected() {
        return store.getters["scribble/getWordSelected"];
    }
}
