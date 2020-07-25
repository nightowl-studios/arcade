import CookieService from "@/services/cookieService";
import { store } from "@/store";

// Event handler for Chat API
export default class AuthHandler {
    handle(payload) {
        const cookieService = new CookieService();
        cookieService.setArcadeCookie(payload.tokenMessage);
        store.commit("application/setPlayerUuid", payload.uuid);
    }
}
