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
    getGameState: (state) => {
        return state.gameState;
    },
    getWordSelected: (state) => {
        return state.wordSelected;
    },
    getRoundNumber: (state) => {
        return state.roundNumber;
    },
    getLoading: (state) => {
        return state.loading;
    },
    getPlayers: (state) => {
        return state.players;
    },
    getPlayerWithUuid: (state) => {
        console.log("getting player with uuid");
        return (uuid) => state.players.filter((p) => p.uuid === uuid)[0];
    },
    getPlayerUuid: (state) => {
        return state.player.uuid;
    },
    getPlayerNickname: (state) => {
        return state.player.nickname;
    },
}

const mutations = {
    setGameState: (state, payload) => {
        state.gameState = payload;
    },
    setWordSelected: (state, payload) => {
        state.wordSelected = payload;
    },
    setRoundNumber: (state, payload) => {
        state.roundNumber = payload;
    },
    setLoading: (state, payload) => {
        state.loading = payload;
    },
    setPlayers: (state, payload) => {
        state.players = payload;
    },
    setPlayer: (state, payload) => {
        state.player = payload;
    },
    setPlayerUuid: (state, payload) => {
        state.player.uuid = payload;
    },
    setPlayerNickname: (state, payload) => {
        state.player.nickname = payload;
    },
    setPlayerReadyState: (state, payload) => {
        const playerUuid = payload.clientUUID;
        const isReady = payload.isReady;

        for (let index = 0; index < state.players.length; index++) {
            if (state.players[index].uuid === playerUuid) {
                state.players[index].isReady = isReady;
                break;
            }
        }
    },
    setPlayerScore: (state, payload) => {
        const playerUuid = payload.uuid;
        const score = payload.score;

        for (let index = 0; index < state.players.length; index++) {
            if (state.players[index].uuid === playerUuid) {
                state.players[index].score = score;
                break;
            }
        }
    },
    addPlayer: (state, payload) => {
        state.players.push(payload);
    },
    removePlayer: (state, payload) => {
        const index = state.players.indexOf(payload);
        state.players.splice(index, 1);
    },
    updateNickname: (state, payload) => {
        const uuid = payload.uuid;
        const nickname = payload.nickname;
        const playerToUpdate = state.players.filter((p) => p.uuid === uuid)[0];
        playerToUpdate.nickname = nickname;
    },
    resetState: (state) => {
        Object.assign(state, getDefaultState());
    }
}

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions: {},
};
