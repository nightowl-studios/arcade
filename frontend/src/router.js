import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";
import Lobby from "./views/Lobby.vue";
import PageNotFound from "./views/PageNotFound.vue";
import Scribble from "./views/Scribble.vue";

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: "/",
            name: "home",
            component: Home
        },
        {
            path: "/lobby",
            redirect: "/"
        },
        {
            path: "/lobby/:lobbyId",
            name: "lobby",
            component: Lobby,
        },
        {
            path: "/scribble",
            name: "scribble",
            component: Scribble,
        },
        {
            path: "/404",
            name: "404",
            component: PageNotFound
        },
        {
            path: "*",
            redirect: "/404"
        }
    ]
})
