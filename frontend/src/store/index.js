import Vue from 'vue'
import Vuex from 'vuex'
//import createLogger from '../../../src/plugins/logger'
import application from './modules/application'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  modules: {
    application
  },
  strict: debug,
  //plugins: debug ? [createLogger()] : []
})
