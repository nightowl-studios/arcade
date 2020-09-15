import Vue from "vue";
import VueCookies from "vue-cookies";

Vue.use(VueCookies);

// Service for handling browser cookies.
export default class CookieService {
    constructor() { }

    getArcadeCookie() {
        return Vue.$cookies.get("arcade_session");
    }

    setArcadeCookie(data) {
        Vue.$cookies.set("arcade_session", data);
    }
}
