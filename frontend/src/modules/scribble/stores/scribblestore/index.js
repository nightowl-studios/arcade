export default {
    namespaced: true,
    state: {
        gameState: null
    },
    getters: {
    },
    mutations: {
        setGameState: (state, payload) => {
            state.gameState = payload
        }
    },
    actions: {

    }
};
