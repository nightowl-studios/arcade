export default {
    namespaced: true,
    state: {
        chosenUuid: "",
        isCanvasLocked: true
    },
    getters: {
    },
    mutations: {
        setChosenUuid: (state, payload) => {
            state.chosenUuid = payload;
        },
        setIsCanvasLocked: (state, payload) => {
            state.isCanvasLocked = payload
        }
    },
    actions: {

    }
};
