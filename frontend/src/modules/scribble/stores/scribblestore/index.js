export default {
    namespaced: true,
    state: {
        gameState: null,
    },
    getters: {
        getGameState: (state) => {
            return state.gameState;
        },
    },
    mutations: {
        setGameState: (state, payload) => {
            state.gameState = payload;
        },
    },
    actions: {},
};
