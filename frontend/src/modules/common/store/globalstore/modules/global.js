const state = {
    message: "Hello from vuex"
}

const getters = {
    getMessage: (state) => {
        return state.message;
    }
}

export default {
    state,
    getters
}
