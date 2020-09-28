const getDefaultState = () => {
    return {
        lobbyId: "",
        tempNickname: null
    }
}

const state = getDefaultState();

const getters = {
    getLobbyId: (storeState) => {
        return storeState.lobbyId;
    },
    getNickname: (storeState) => {
        return storeState.tempNickname;
    }
}

const mutations = {
    setLobbyId: (storeState, payload) => {
        storeState.lobbyId = payload;
    },
    setNickname: (storeState, payload) => {
        storeState.tempNickname = payload;
    },
    resetState: (storeState) => {
        Object.assign(storeState, getDefaultState());
    }
}

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions: {},
};
