import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import Vue from 'vue'
import App from './App.vue'
import './index.scss'
import store from './modules/common/store/globalstore/index'
import router from './router'
import CookieService from './services/cookieService'
import EventHandlerService from './services/eventHandlerService'
import HubApiService from './services/hubApiService'
import WebSocketService from './services/webSocketService'

Vue.config.productionTip = false

// Install BootstrapVue
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)

// Global Instance Properties
Vue.prototype.$hubAPI = 'hub'

let webSocketURL = 'ws://' + document.location.hostname + ':8081/ws'
let httpURL = 'http://' + document.location.hostname + ':8081'
Vue.prototype.$cookieService = new CookieService()
let eventHandlerService = new EventHandlerService()
Vue.prototype.$webSocketService = new WebSocketService(
    webSocketURL,
    Vue.prototype.$cookieService,
    eventHandlerService
)
Vue.prototype.$hubApiService = new HubApiService(httpURL)

new Vue({
    store,
    router,
    render: (h) => h(App),
}).$mount('#app')
