export default {
    namespaced: true,
    state: {
        gameState: null,
        wordSelected: "",
    },
    getters: {
        getGameState: (state) => {
            return state.gameState;
        },
        getWordSelected: (state) => {
            return state.wordSelected;
        },
    },
    mutations: {
        setGameState: (state, payload) => {
            state.gameState = payload;
        },
        setWordSelected: (state, payload) => {
            state.wordSelected = payload;
        },
    },
    actions: {},
};
