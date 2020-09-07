import { EventBus } from "@/eventBus";
import { Event } from "@/events.js";
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

            this.handlePlayersChanged(players);

            store.commit("application/setPlayers", players);
        }
    }

    handlePlayersChanged(players) {
        const currentPlayers = store.getters["application/getPlayers"];
        if (currentPlayers.length !== 0) {
            const playersJoined = this.getPlayerListDifference(
                players,
                currentPlayers
            );
            playersJoined.forEach((player) => {
                EventBus.$emit(Event.PLAYER_JOIN, player);
            });

            const playersLeft = this.getPlayerListDifference(
                currentPlayers,
                players
            );
            playersLeft.forEach((player) => {
                EventBus.$emit(Event.PLAYER_LEFT, player);
            });
        }
    }

    getPlayerListDifference(playerList1, playerList2) {
        return playerList1.filter(
            (p) => !playerList2.map((q) => q.uuid).includes(p.uuid)
        );
    }
}
