<template>
  <div>
    <LobbyText :clients="clients" />
    <div>{{connectionState}}</div>
    <div>Room Id: {{ lobbyId }}</div>
    <Nickname @onChangeNickname="onChangeNickname" />
    <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
    <b-button v-on:click="goToScribble()">Go to Scribble</b-button>
  </div>
</template>

<script>
import LobbyText from "../components/LobbyText.vue";
import Nickname from "../components/Nickname.vue";
import { EventBus } from "../eventBus.js";

export default {
  name: "Lobby",
  components: {
    LobbyText,
    Nickname
  },
  data: function() {
    return {
      clients: [],
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
      this.$router.push({ path: "/scribble" });
    }
  },
  async created() {
    EventBus.$on("connected", () => {
      this.connectionState = "CONNECTED";
    }),
      EventBus.$on(this.$hubAPI, data => {
        this.clients = data.connectedClients;
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
