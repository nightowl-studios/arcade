import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        message: "HELLO FROM VUEX",
    },
    mutations: {
        setMessage(state, payload) {
            state.message = payload;
        }
    },
    actions: {

    },
    getters: {
        message(state) {
            return state.message;
        }
    }
});