const getDefaultState = () => {
    return {
        loading: true,
        gameState: null,
        wordSelected: "",
        roundNumber: 0,
        players: [],
        player: null,
    }
}

const state = getDefaultState();

const getters = {
    getGameState: (storeState) => {
        return storeState.gameState;
    },
    getWordSelected: (storeState) => {
        return storeState.wordSelected;
    },
    getRoundNumber: (storeState) => {
        return storeState.roundNumber;
    },
    getLoading: (storeState) => {
        return storeState.loading;
    },
    getPlayers: (storeState) => {
        return storeState.players;
    },
    getPlayerWithUuid: (storeState) => {
        return (uuid) => storeState.players.filter((p) => p.uuid === uuid)[0];
    },
    getPlayerUuid: (storeState) => {
        return storeState.player.uuid;
    },
    getPlayerNickname: (storeState) => {
        return storeState.player.nickname;
    },
}

const mutations = {
    setGameState: (storeState, payload) => {
        storeState.gameState = payload;
    },
    setWordSelected: (storeState, payload) => {
        storeState.wordSelected = payload;
    },
    setRoundNumber: (storeState, payload) => {
        storeState.roundNumber = payload;
    },
    setLoading: (storeState, payload) => {
        storeState.loading = payload;
    },
    setPlayers: (storeState, payload) => {
        storeState.players = payload;
    },
    setPlayer: (storeState, payload) => {
        storeState.player = payload;
    },
    setPlayerUuid: (storeState, payload) => {
        storeState.player.uuid = payload;
    },
    setPlayerNickname: (storeState, payload) => {
        storeState.player.nickname = payload;
    },
    setPlayerReadyState: (storeState, payload) => {
        const playerUuid = payload.clientUUID;
        const isReady = payload.isReady;

        for (let index = 0; index < storeState.players.length; index++) {
            if (storeState.players[index].uuid === playerUuid) {
                storeState.players[index].isReady = isReady;
                break;
            }
        }
    },
    setPlayerScore: (storeState, payload) => {
        const playerUuid = payload.uuid;
        const score = payload.score;

        for (let index = 0; index < storeState.players.length; index++) {
            if (storeState.players[index].uuid === playerUuid) {
                storeState.players[index].score = score;
                break;
            }
        }
    },
    addPlayer: (storeState, payload) => {
        storeState.players.push(payload);
    },
    removePlayer: (storeState, payload) => {
        const index = storeState.players.indexOf(payload);
        storeState.players.splice(index, 1);
    },
    updateNickname: (storeState, payload) => {
        const uuid = payload.uuid;
        const nickname = payload.nickname;
        const playerToUpdate = storeState.players.filter((p) => p.uuid === uuid)[0];
        playerToUpdate.nickname = nickname;
    },
    resetState: (storeState) => {
        Object.assign(storeState, getDefaultState());
    }
}

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions: {},
};
