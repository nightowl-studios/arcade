import Vue from "vue";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import App from "./App.vue";
import router from "./router";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "./index.scss";
import store from './store/index'
import WebSocketService from './services/webSocketService'
import CookieService from "./services/cookieService";

Vue.config.productionTip = false;

// Install BootstrapVue
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

// Global Instance Properties
Vue.prototype.$httpURL = "http://" + document.location.hostname +":8081";
Vue.prototype.$hubAPI = "hub";

let webSocketURL = "ws://" + document.location.hostname +":8081/ws";
Vue.prototype.$cookieService = new CookieService();
Vue.prototype.$webSocketService = new WebSocketService(webSocketURL, Vue.prototype.$cookieService);

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount("#app");
