import ApiReceiverService from "@/backend/apiservice/receive/apiReceiverService";
import ApiSenderFacade from "@/backend/apiservice/send/apiSenderFacade";
import ChatApiService from "@/backend/apiservice/send/chatApiService";
import DrawApiService from "@/backend/apiservice/send/drawApiService";
import GameApiService from "@/backend/apiservice/send/gameApiService";
import HubApiService from "@/backend/apiservice/send/hubApiService";
import WebSocketConnection from "@/backend/communication/webSocketConnection";
import WebSocketService from "@/backend/communication/webSocketService";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Vue from "vue";
import VueSimpleAlert from "vue-simple-alert";
import App from "./App.vue";
import AuthApiService from "./backend/apiservice/send/authApiService";
import ApplicationController from "./backend/application/applicationController";
import ApplicationStoreService from "./backend/application/applicationStoreService";
import CookieService from "./backend/application/cookieService";
import GameManager from "./backend/scribble/gameManager";
import ScribbleGameController from "./backend/scribble/scribbleGameController";
import ScribbleReceiver from "./backend/scribble/scribbleReceiver";
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

const webSocketConnection = new WebSocketConnection();

const webSocketService = new WebSocketService(webSocketUrl, webSocketConnection);
apiSenderFacade.setWebSocketService(webSocketService);
const hubApiService = new HubApiService(httpUrl, webSocketService);
apiSenderFacade.setHubApiService(hubApiService);
const gameApiService = new GameApiService(webSocketService);
apiSenderFacade.setGameApiService(gameApiService);
const chatApiService = new ChatApiService(webSocketService);
apiSenderFacade.setChatApiService(chatApiService);
const drawApiService = new DrawApiService(webSocketService);
apiSenderFacade.setDrawApiService(drawApiService);
const cookieService = new CookieService();
const authApiService = new AuthApiService(webSocketService, cookieService);
apiSenderFacade.setAuthApiService(authApiService);
Object.freeze(apiSenderFacade);

Vue.prototype.$scribbleStoreService = new ScribbleStoreService();
Vue.prototype.$scribbleGameController = new ScribbleGameController(apiSenderFacade, Vue.prototype.$scribbleStoreService);

Vue.prototype.$applicationStoreService = new ApplicationStoreService();
Vue.prototype.$applicationController = new ApplicationController(apiSenderFacade, Vue.prototype.$applicationStoreService);

const apiReceiverService = new ApiReceiverService();

const scribbleGameManager = new GameManager(Vue.prototype.$scribbleGameController, Vue.prototype.$applicationStoreService, Vue.prototype.$scribbleStoreService)
const scribbleReceiver = new ScribbleReceiver(scribbleGameManager);

webSocketService.getConnection().addListener(apiReceiverService);
apiReceiverService.addListener(scribbleReceiver);


new Vue({
    store,
    router,
    render: (h) => h(App),
}).$mount("#app");
