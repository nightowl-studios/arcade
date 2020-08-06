import Vue from "vue";
import Router from "vue-router";
import PageNotFound from "./modules/common/view/PageNotFound.vue";
import Home from "./modules/home/view/Home.vue";
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
