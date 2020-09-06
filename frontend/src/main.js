import ApiReceiverService from "@/backend/apiservice/receive/apiReceiverService";
import ApiSenderFacade from "@/backend/apiservice/send/apiSenderFacade";
import ChatApiService from "@/backend/apiservice/send/chatApiService";
import DrawApiService from "@/backend/apiservice/send/drawApiService";
import GameApiService from "@/backend/apiservice/send/gameApiService";
import HubApiService from "@/backend/apiservice/send/hubApiService";
import WebSocketService from "@/backend/communication/webSocketService";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Vue from "vue";
import VueSimpleAlert from "vue-simple-alert";
import App from "./App.vue";
import ApplicationController from "./backend/application/applicationController";
import ApplicationStoreService from "./backend/application/applicationStoreService";
import ScribbleGameController from "./backend/scribble/scribbleGameController";
import ScribbleStoreService from "./backend/scribble/scribbleStoreService";
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

const apiSenderFacade = ApiSenderFacade;
const webSocketService = new WebSocketService(webSocketUrl);
apiSenderFacade.setWebSocketService(webSocketService);
const hubApiService = new HubApiService(httpUrl);
apiSenderFacade.setHubApiService(hubApiService);
const gameApiService = new GameApiService();
apiSenderFacade.setGameApiService(gameApiService);
const chatApiService = new ChatApiService();
apiSenderFacade.setChatApiService(chatApiService);
const drawApiService = new DrawApiService();
apiSenderFacade.setDrawApiService(drawApiService);
Object.freeze(apiSenderFacade);

const scribbleStoreService = new ScribbleStoreService();
Vue.prototype.$scribbleGameController = new ScribbleGameController(apiSenderFacade, scribbleStoreService);

const applicationStoreService = new ApplicationStoreService();
Vue.prototype.$applicationController = new ApplicationController(apiSenderFacade, applicationStoreService);

const apiReceiverService = new ApiReceiverService(
    Vue.prototype.$applicationController,
    Vue.prototype.$scribbleGameController,
    applicationStoreService,
    scribbleStoreService);
Object.freeze(apiReceiverService);

webSocketService.getConnection().addListener(apiReceiverService);



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
