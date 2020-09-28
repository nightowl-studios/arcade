const getDefaultState = () => {
    return {
        lobbyId: "",
        tempNickname: null
    }
}

const state = getDefaultState();

const getters = {
    getLobbyId: (state) => {
        return state.lobbyId;
    },
    getNickname: (state) => {
        return state.tempNickname;
    }
}

const mutations = {
    setLobbyId: (state, payload) => {
        state.lobbyId = payload;
    },
    setNickname: (state, payload) => {
        state.tempNickname = payload;
    },
    resetState: (state) => {
        Object.assign(state, getDefaultState());
    }
}

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions: {},
};
