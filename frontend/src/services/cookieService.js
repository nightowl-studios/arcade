import Vue from 'vue';
import VueCookies from 'vue-cookies'

Vue.use(VueCookies);

export default class CookieService {
    constructor() {}

    getArcadeCookie() {
        console.log("Getting arcade cookie");
        return Vue.$cookies.get('arcade_session');
    }

    setArcadeCookie(data) {
        console.log("Setting arcade cookie");
        Vue.$cookies.set("arcade_session", data);
    }
}