<template>
  <div id="app">
    <div v-if="connectionState === 'CONNECTED'">
      <Lobby/>
    </div>
    <div v-else>
      <img alt="Vue logo" src="./assets/logo.png">
      <HelloWorld msg="Welcome to Your Vue.js App"/>
      <CreateButton @onCreateRoom="onCreateRoom"/>
      <b-button v-on:click="sendMessage('hello')">Send a Message</b-button>
      <JoinModal @onJoinRoom="onJoinRoom"/>
      <div>{{connectionState}} : {{hubId}}</div>
    </div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Lobby from './components/Lobby.vue'
import CreateButton from './components/CreateButton.vue'
import JoinModal from './components/JoinModal.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    Lobby,
    CreateButton,
    JoinModal
  },
  data: function() {
    return {
      connection: null,
      isConnected: false,
      players: [
        { name: "Gordon", id: "ID12345"},
        { name: "Byron", id: "ID12346"},
        { name: "Zach", id: "ID12347" },
        { name: "Sam", id: "ID12348" }
      ],
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

      this.connection.onmessage = function(event) {
        console.log(event);
      }

      this.connection.onopen = (event) => {
        console.log(event);
        console.log("Successfully connected to the websocket...");
        console.log(this);
        this.connectionState = "CONNECTED";
      }
    },
<<<<<<< Updated upstream
=======
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
    
>>>>>>> Stashed changes
    sendMessage: function(message) {
      message = {
        "api":"echo",
        "payload":{
          "message":"zacsdfsdfsdfsdfary"
        }
      }
      let json = JSON.stringify(message);
      console.log(json)
      console.log(this.connection);
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

        this.connection.onmessage = function(event) {
          console.log(event);
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
