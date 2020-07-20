import { GlobalStore } from "../store/globalstore/index";

export default {
    computed: {
        players() {
            return GlobalStore.getters.getPlayers;
        },
    },
    created: async function () {
        this.lobbyId = this.$router.currentRoute.params.lobbyId;
        if (!this.$webSocketService.isConnected()) {
            let lobbyExists = await this.$hubApiService.checkLobbyExists(
                this.lobbyId
            );
            if (!lobbyExists) {
                this.$router.push({ name: "404" });
            } else {
                this.$webSocketService.connect(this.lobbyId);
            }
        }
    },
};
