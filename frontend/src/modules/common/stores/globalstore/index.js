export default {
    namespaced: true,
    state: {
        players: []
    },
    getters: {
    },
    mutations: {
        setPlayers: (state, payload) => {
            state.players = payload;
        },
    },
    actions: {

    }
};
