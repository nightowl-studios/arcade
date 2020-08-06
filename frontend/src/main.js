import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import VueSimpleAlert from "vue-simple-alert";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Vue from "vue";
import App from "./App.vue";
import "./index.scss";
import router from "./router";
import ChatApiService from "./services/chatApiService";
import CookieService from "./services/cookieService";
import EventHandlerService from "./services/eventHandlerService";
import GameApiService from "./services/gameApiService";
import HubApiService from "./services/hubApiService";
import WebSocketService from "./services/webSocketService";
import { store } from "./store";

Vue.config.productionTip = false;

// Install BootstrapVue
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);
// Install VueSimpleAlert
Vue.use(VueSimpleAlert);

// Global Instance Properties
Vue.prototype.$hubAPI = "hub";

const webSocketURL = `ws://${document.location.hostname}:8081/ws`;
const httpURL = `http://${document.location.hostname}:8081`;

const cookieService = new CookieService();
const eventHandlerService = new EventHandlerService();
Vue.prototype.$webSocketService = new WebSocketService(
    webSocketURL,
    cookieService,
    eventHandlerService
);

// API Services
Vue.prototype.$hubApiService = new HubApiService(httpURL);
Vue.prototype.$gameApiService = new GameApiService(
    Vue.prototype.$webSocketService
);
Vue.prototype.$chatApiService = new ChatApiService(
    Vue.prototype.$webSocketService
);

new Vue({
    store,
    router,
    render: (h) => h(App),
}).$mount("#app");
