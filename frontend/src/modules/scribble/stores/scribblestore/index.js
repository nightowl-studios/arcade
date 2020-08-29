export default {
    namespaced: true,
    state: {
        loading: true,
        gameState: null,
        wordSelected: "",
        roundNumber: 0,
        players: [],
    },
    getters: {
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
            return (uuid) => state.players.filter((p) => p.uuid === uuid)[0];
        },
    },
    mutations: {
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
    },
    actions: {},
};
