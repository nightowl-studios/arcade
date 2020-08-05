export default {
    namespaced: true,
    state: {
        playerUuid: "",
        players: [],
    },
    getters: {
        getPlayerUuid: (state) => {
            return state.playerUuid;
        },
        getPlayerWithUuid: (state) => {
            return (uuid) => state.players.filter((p) => p.uuid === uuid)[0];
        },
    },
    mutations: {
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
        setPlayerUuid: (state, payload) => {
            state.playerUuid = payload;
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
