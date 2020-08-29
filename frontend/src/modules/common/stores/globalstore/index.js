export default {
    namespaced: true,
    state: {
        lobbyId: "",
        playerUuid: "",
        players: [],
        nickname: "",
    },
    getters: {
        getLobbyId: (state) => {
            return state.lobbyId;
        },
        getPlayers: (state) => {
            return state.players;
        },
        getPlayerUuid: (state) => {
            return state.playerUuid;
        },
        getPlayerWithUuid: (state) => {
            return (uuid) => state.players.filter((p) => p.uuid === uuid)[0];
        },
        getNickname: (state) => {
            return state.nickname;
        },
    },
    mutations: {
        setLobbyId: (state, payload) => {
            state.lobbyId = payload;
        },
        setNickname: (state, payload) => {
            state.nickname = payload;
        },
        setPlayerUuid: (state, payload) => {
            state.playerUuid = payload;
        },
        setPlayers: (state, payload) => {
            state.players = payload;
        },
        setPlayerReadyState: (state, payload) => {
            const playerUuid = payload.clientUUID;
            const isReady = payload.isReady;

            for (let index = 0; index < state.players.length; index++) {
                if (state.players[index].uuid === playerUuid) {
                    state.players[index].isReady = isReady;
                    break;
                }
            }
        },
        setPlayerScore: (state, payload) => {
            const playerUuid = payload.uuid;
            const score = payload.score;

            for (let index = 0; index < state.players.length; index++) {
                if (state.players[index].uuid === playerUuid) {
                    state.players[index].score = score;
                    break;
                }
            }
        },
    },
    actions: {},
};
