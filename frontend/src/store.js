import application from "@/modules/common/stores/globalstore/index.js";
import scribble from "@/modules/scribble/stores/scribblestore/index.js";
import Vue from "vue";
import Vuex from "vuex";


Vue.use(Vuex);

const debug = process.env.NODE_ENV !== "production";

export const store = new Vuex.Store({
    modules: {
        application,
        scribble
    },
    strict: debug,
})
