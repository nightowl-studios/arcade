<template>
  <div id="app">
    <div v-if="connectionState === 'CONNECTED'">
      <Lobby :clients="clients" />
      <div>{{connectionState}}</div>
      <div>Room Id: {{ hubId }}</div>
      <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
    </div>
    <div v-else>
      <Title msg="Not ScribbleIO" />
      <HelloWorld msg="Welcome to Your Vue.js App" />
      <CreateButton @onCreateRoom="onCreateRoom" />
      <JoinModal @onJoinRoom="onJoinRoom" />
      <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
      <CanvasPanel :colors="colors" :sizes="sizes" />
    </div>
  </div>
</template>

<script>
import Title from "./components/Title.vue";
import Lobby from "./components/Lobby.vue";
import CreateButton from "./components/CreateButton.vue";
import JoinModal from "./components/JoinModal.vue";
import CanvasPanel from "./components/CanvasPanel.vue";
import { EventBus } from "./eventBus.js";
import { ArcadeWebSocket } from "./webSocket.js";

export default {
  name: "App",
  components: {
    Title,
    Lobby,
    CreateButton,
    JoinModal,
    CanvasPanel
  },
  data: function() {
    return {
      connection: null,
      clients: [],
      lobbyId: "",
      connectionState: "DISCONNECTED",
      colors: ["#000000", "#4287f5", "#da42f5", "#7ef542"],
      sizes: [8, 16, 32, 64]
    };
  },
  methods: {
    onCreateRoom: function(lobbyId) {
      this.hubId = lobbyId;
    },
    onJoinRoom: function(lobbyId) {
      this.hubId = lobbyId;
    },
    sendPlayerMessage: function() {
      let message = {
        api: "hub",
        payload: {
          requestLobbyDetails: true
        }
      };
      ArcadeWebSocket.send(message);
    }
  },
  created() {
    EventBus.$on("connected", () => {
      this.connectionState = "CONNECTED";
    }),
      EventBus.$on(this.$hubAPI, data => {
        this.clients = data.connectedClients;
      });
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
