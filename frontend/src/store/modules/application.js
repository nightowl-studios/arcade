const state = () => ({
  message: "Hello from vuex",
});

const getters = {
  getMessage: (state) => {
    return state.message;
  },
};

const actions = {};

const mutations = {};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
