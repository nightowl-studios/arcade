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

    setPlayerScore(playerScore) {
        store.commit("scribble/setPlayerScore", playerScore);
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

    setPlayer(player) {
        store.commit("scribble/setPlayer", player);
    }

    setPlayerUuid(uuid) {
        store.commit("scribble/setPlayerUuid", uuid);
    }

    getPlayerUuid() {
        return store.getters["scribble/getPlayerUuid"];
    }

    setNickname(nickname) {
        store.commit("scribble/setPlayerNickname", nickname);
    }

    getNickname() {
        return store.getters["scribble/getNickname"];
    }

    getTimerDuration() {
        const state = this.getState();
        return state.duration;
    }

    addPlayer(player) {
        store.commit("scribble/addPlayer", player);
    }

    removePlayer(player) {
        store.commit("scribble/removePlayer", player);
    }

    updateNickname(uuid, nickname) {
        const payload = {
            uuid: uuid,
            nickname: nickname
        };
        store.commit("scribble/updateNickname", payload);
    }

    setScore(playerUuid, score) {
        const payload = {
            playerUuid: playerUuid,
            score: score
        };
        store.commit("scribble/setScore", payload);
    }
}
