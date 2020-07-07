<template>
  <div id="app">
    <div v-if="connectionState === 'CONNECTED'">
      <Lobby :clients="clients"/>
      <div>{{connectionState}}</div>
      <div>Room Id: {{ hubId }}</div>
      <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
    </div>
    <div v-else>
      <Title msg="Not ScribbleIO"/>
      <HelloWorld msg="Welcome to Your Vue.js App"/>
      <CreateButton @onCreateRoom="onCreateRoom"/>
      <JoinModal @onJoinRoom="onJoinRoom"/>
      <Canvas/>
    </div>
  </div>
</template>

<script>
import Title from './components/Title.vue'
import Lobby from './components/Lobby.vue'
import CreateButton from './components/CreateButton.vue'
import JoinModal from './components/JoinModal.vue'
import Canvas from './components/Canvas.vue'
import { EventBus } from './eventBus.js';

export default {
  name: 'App',
  components: {
    Title,
    Lobby,
    CreateButton,
    JoinModal,
    Canvas
  },
  data: function() {
    return {
      connection: null,
      clients: [],
      hubId: "",
      connectionState: "DISCONNECTED"
    }
  },
  methods: {
    onCreateRoom: function(lobbyId) {
      console.log(lobbyId);
      console.log("Connecting to websocket...");
      this.connectionState = "CONNECTING";

      this.hubId = lobbyId;
      EventBus.connect(lobbyId);
    },
    onJoinRoom: function(event) {
      console.log("Checking if hubId exists...");

      if (event.response.data.exists) {
        console.log("Connecting to websocket...");
        this.connectionState = "CONNECTING";
        this.hubId = event.hubId;
        EventBus.connect(event.hubId);
      } else {
        console.log("HubId does not exist...");
      }
    },
    sendPlayerMessage: function(){
      let message = {
        "api":"hub",
        "payload":{
          "requestLobbyDetails":true
        }
      }
      let json = JSON.stringify(message);
      this.connection.send(json);
    }
  },
  created() {
    EventBus.$on('connected', () => {
      this.connectionState = "CONNECTED";
    }),
    EventBus.$on(this.$hubAPI, (data) => {
      this.clients = data.connectedClients;
    }) 
  }
}
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
