import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Vue from "vue";
import App from "./App.vue";
import "./index.scss";
import router from "./router";
import store from "./store/index";

Vue.config.productionTip = false;

// Install BootstrapVue
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

// Global Instance Properties
Vue.prototype.$httpURL = "http://" + document.location.hostname +":8081";
Vue.prototype.$websocketURL = "ws://" + document.location.hostname +":8081/ws";
Vue.prototype.$hubAPI = "hub";

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount("#app");
