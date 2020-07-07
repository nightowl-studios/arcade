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
    onCreateRoom: function(event) {
      console.log("Connecting to websocket...");
      this.connectionState = "CONNECTING";
      this.hubId = event.data.hubID;

      let webSocketUrl = this.$websocketURL + "/" + this.hubId;
      this.connection = new WebSocket(webSocketUrl);

      this.connection.onmessage = (event) => {
        console.log(event.data);
        let parseMsg = JSON.parse(event.data);
        console.log(parseMsg);
        this.clients = parseMsg.payload.connectedClients;
        console.log(parseMsg.payload.connectedClients[0].clientUUID)
      }

      this.connection.onopen = (event) => {
        console.log(event);
        console.log("Successfully connected to the websocket...");
        console.log(this);
        this.connectionState = "CONNECTED";
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
    },
    onJoinRoom: function(event) {
      console.log("Checking if hubId exists...");

      console.log(event);
      if (event.response.data.exists) {
        console.log("Connecting to websocket...");
        this.connectionState = "CONNECTING";
        this.hubId = event.hubId;
        console.log(event);

        let webSocketUrl = this.$websocketURL + "/" + this.hubId;
        this.connection = new WebSocket(webSocketUrl);

        this.connection.onmessage = (event) => {
          console.log(event.data);
          let parseMsg = JSON.parse(event.data);
          console.log(parseMsg);
          this.clients = parseMsg.payload.connectedClients;
          console.log(parseMsg.payload.connectedClients[0].clientUUID)
        }

        this.connection.onopen = (event) => {
          console.log(event);
          console.log("Successfully connected to the websocket...");
          console.log(this);
          this.connectionState = "CONNECTED";
        }
      } else {
        console.log("HubId does not exist...");
      }
    }
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
