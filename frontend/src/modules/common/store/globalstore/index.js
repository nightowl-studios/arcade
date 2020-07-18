import Vue from 'vue'
import Vuex from 'vuex'
import global from './modules/global'

Vue.use(Vuex)

export const store = new Vuex.Store({
  modules: {
    global
  }
})
