import ApiSenderFacade from "@/backend/apiservice/send/apiSenderFacade";
import GameApiService from "@/backend/apiservice/send/gameApiService";
import HubApiService from "@/backend/apiservice/send/hubApiService";
import WebSocketService from "@/backend/communication/webSocketService";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Vue from "vue";
import VueSimpleAlert from "vue-simple-alert";
import App from "./App.vue";
import "./index.scss";
import router from "./router";
import { store } from "./store";



Vue.config.productionTip = false;

// Install BootstrapVue
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);
// Install VueSimpleAlert
Vue.use(VueSimpleAlert);

const webSocketUrl = `ws://${document.location.hostname}:8081/ws`;
const httpUrl = `http://${document.location.hostname}:8081`;

Vue.prototype.$webSocketService = new WebSocketService(webSocketUrl);

const hubApiService = new HubApiService(httpUrl);
const gameApiService = new GameApiService();
Vue.prototype.$apiSenderFacade = ApiSenderFacade;
ApiSenderFacade.setWebSocketService(Vue.prototype.$webSocketService);
ApiSenderFacade.setHubApiService(hubApiService);
ApiSenderFacade.setGameApiService(gameApiService);
Object.freeze(ApiSenderFacade);


/* DEPRECATED CODE */
// Global Instance Properties
// Vue.prototype.$hubAPI = "hub";

// const webSocketURL = `ws://${document.location.hostname}:8081/ws`;
// const httpURL = `http://${document.location.hostname}:8081`;

// const cookieService = new CookieService();
// const eventHandlerService = new EventHandlerService();
// Vue.prototype.$webSocketService = new WebSocketService(
//     webSocketURL,
//     cookieService,
//     eventHandlerService
// );

// // API Services
// Vue.prototype.$hubApiService = new HubApiService(httpURL);
// Vue.prototype.$gameApiService = new GameApiService(
//     Vue.prototype.$webSocketService
// );
// Vue.prototype.$chatApiService = new ChatApiService(
//     Vue.prototype.$webSocketService
// );
/* END OF DEPRECATED CODE */

new Vue({
    store,
    router,
    render: (h) => h(App),
}).$mount("#app");
