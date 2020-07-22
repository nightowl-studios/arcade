export default {
    namespaced: true,
    state: {
        chosenUuid: ""
    },
    getters: {
    },
    mutations: {
        setChosenUuid: (state, payload) => {
            state.chosenUuid = payload;
        },
    },
    actions: {

    }
};
