export default {
    namespaced: true,
    state: {
        playerUuid: "",
        players: []
    },
    getters: {
    },
    mutations: {
        setPlayers: (state, payload) => {
            state.players = payload;
        },
        setPlayerUuid: (state, payload) => {
            state.playerUuid = payload;
        }
    },
    actions: {

    }
};
