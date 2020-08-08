export default {
    namespaced: true,
    state: {
        gameState: null,
        wordSelected: "",
        roundNumber: 0,
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
    },
    actions: {},
};
