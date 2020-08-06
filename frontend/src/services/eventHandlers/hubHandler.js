import { EventBus } from "@/eventBus";
import { Event } from "@/events";
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

            this.handlePlayerJoined(players);
            this.handlePlayerLeft(players);

            store.commit("application/setPlayers", players);
        }
    }

    handlePlayerJoined(players) {
        const currentPlayers = store.getters["application/getPlayers"];
        if (currentPlayers.length !== 0) {
            const clientUuids = players.map((player) => player.uuid);
            const currentUuids = currentPlayers.map((player) => player.uuid);

            let playerJoinGame = clientUuids.filter(
                (x) => !currentUuids.includes(x)
            );

            if (playerJoinGame.length === 1) {
                const player = players.filter(
                    (p) => p.uuid === playerJoinGame[0]
                )[0];
                EventBus.$emit(Event.PLAYER_JOIN, player);
            }
        }
    }

    handlePlayerLeft(players) {
        const currentPlayers = store.getters["application/getPlayers"];
        if (currentPlayers.length !== 0) {
            const clientUuids = players.map((player) => player.uuid);
            const currentUuids = currentPlayers.map((player) => player.uuid);

            let playerLeftGame = currentUuids.filter(
                (x) => !clientUuids.includes(x)
            );

            if (playerLeftGame.length === 1) {
                const player = store.getters["application/getPlayerWithUuid"](
                    playerLeftGame[0]
                );
                EventBus.$emit(Event.PLAYER_LEFT, player);
            }
        }
    }
}
