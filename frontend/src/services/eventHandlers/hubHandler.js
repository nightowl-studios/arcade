import { store } from "@/store";

// Event handler for Hub API
export default class HubHandler {
    handle(payload) {
        if (payload.connectedClients != null) {
            store.commit("application/setPlayers", payload.connectedClients);
        }
    }
}
