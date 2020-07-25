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
        setPlayerUuid: (state, payload) => {
            state.playerUuid = payload;
        },
    },
    actions: {},
};
