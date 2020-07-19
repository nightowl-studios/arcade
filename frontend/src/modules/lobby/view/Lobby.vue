<template>
  <div id="lobby">
    <Header title="Welcome to Not Scribble" />
    <Header class="lobby-header-room-id" :title="lobbyId" />
    <div class="lobby-buttons">
      <b-button class="lobby-button" variant="success" v-on:click="goToScribble">Start game</b-button>
      <Nickname @onChangeNickname="onChangeNickname" />
      <b-button class="exit-button" variant="danger" v-on:click="exitToHome">Exit Lobby</b-button>
    </div>
    <PlayerList :players="players" />
  </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import Nickname from "../components/Nickname.vue";
import { EventBus } from "@/eventBus.js";

export default {
  name: "Lobby",
  components: {
    Header,
    PlayerList,
    Nickname
  },
  data: function() {
    return {
      players: [],
      lobbyId: "",
      connectionState: "DISCONNECTED"
    };
  },
  methods: {
    onChangeNickname: function(event) {
      let message = {
        api: "hub",
        payload: {
          changeNameTo: event.nickname
        }
      };
      this.$webSocketService.send(message);
    },
    sendPlayerMessage: function() {
      let message = {
        api: "hub",
        payload: {
          requestLobbyDetails: true
        }
      };
      this.$webSocketService.send(message);
    },
    goToScribble: function() {
      let message = {
        api: "game",
        payload: {
          gameMasterAPI: "waitForStart",
          waitForStart: {
            startGame: true
          }
        }
      };
      this.$webSocketService.send(message);
      this.$router.push({ path: "/scribble" });
    },
    exitToHome: function() {
      this.$webSocketService.disconnect();
      this.$router.push({ name: "home" });
    }
  },
  async created() {
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
</script>
<style scoped>
#lobby {
  padding-left: 100px;
  padding-right: 100px;
}

.lobby-button, .exit-button {
  margin-left: 2px;
  margin-right: 2px;
}

.lobby-header-room-id {
  color: orange;
}
</style>
