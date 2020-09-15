export default {
    namespaced: true,
    state: {
        lobbyId: "",
        tempNickname: null,
    },
    getters: {
        getLobbyId: (state) => {
            return state.lobbyId;
        },
        getNickname: (state) => {
            return state.tempNickname;
        }
    },
    mutations: {
        setLobbyId: (state, payload) => {
            state.lobbyId = payload;
        },
        setNickname: (state, payload) => {
            state.tempNickname = payload;
        }
    },
    actions: {},
};
