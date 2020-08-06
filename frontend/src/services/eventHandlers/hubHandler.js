import Player from "@/modules/common/entities/player";
import { store } from "@/store";

// Event handler for Hub API
export default class HubHandler {
    handle(payload) {
        if (payload.connectedClients != null) {
            const players = payload.connectedClients.map(
                (client) =>
                    new Player(
                        client.clientUUID.UUID,
                        client.nickname,
                        client.joinOrder
                    )
            );
            store.commit("application/setPlayers", players);
        }
    }
}
