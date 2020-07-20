import Vue from "vue";
import Router from "vue-router";
import PageNotFound from "./modules/common/view/PageNotFound.vue";
import Home from "./modules/home/view/Home.vue";
import Lobby from "./modules/lobby/view/Lobby.vue";
import Scribble from "./modules/scribble/view/Scribble.vue";

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: "/",
            name: "home",
            component: Home,
        },
        {
            path: "/lobby",
            redirect: "/",
        },
        {
            path: "/lobby/:lobbyId",
            name: "lobby",
            component: Lobby,
        },
        {
            path: "/scribble",
            redirect: "/",
        },
        {
            path: "/scribble/:lobbyId",
            name: "lobby",
            component: Scribble,
        },
        {
            path: "/404",
            name: "404",
            component: PageNotFound,
        },
        {
            path: "*",
            redirect: "/404",
        },
    ],
});
