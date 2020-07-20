const state = {
    players: [],
};

const getters = {
    getPlayers: (state) => {
        return state.players;
    },
};

const mutations = {
    setPlayers: (state, payload) => {
        state.players = payload;
    },
};

export default {
    state,
    getters,
    mutations,
};
