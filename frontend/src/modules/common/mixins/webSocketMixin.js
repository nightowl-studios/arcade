import { EventBus } from "@/eventBus.js";

export default {
  data: function() {
    return {
      players: []
    };
  },
  created: async function() {
    EventBus.$on("connected", () => {
      this.connectionState = "CONNECTED";
    }),

    EventBus.$on(this.$hubAPI, data => {
      this.players = data.connectedClients;
    });

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
  }
};
