import { GlobalStore } from "@/modules/common/store/globalstore";

// Event handler for Hub API
export default class HubHandler {
    handle(payload) {
        if (payload.connectedClients != null) {
            GlobalStore.commit("setPlayers", payload.connectedClients);
        }
    }
}
