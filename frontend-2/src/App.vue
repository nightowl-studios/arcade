<template>
  <div id="app">
    <div v-if="isConnected">
      <img alt="Vue logo" src="./assets/logo.png">
      <HelloWorld msg="Welcome to Your Vue.js App"/>
      <b-button v-on:click="sendMessage('hello')">Button</b-button>
    </div>
    <div v-else>
      <h1>Lobby</h1>
      <div v-for="player in players" :key="player">
        <Player :name=player.name :id=player.id />
      </div>
    </div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Player from './components/Player.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    Player
  },
  data: function() {
    return {
      connection: null,
      isConnected: false,
      players: [
        { name: "Gordon", id: "something"},
        { name: "Byron", id: "something2"},
        { name: "Zach", id: "something3" },
        { name: "Sam", id: "something4" }
      ]
    }
  },
  methods: {
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
    }
  },
  created: function() {
    console.log("Starting connection to WebSocket Server")
    this.connection = new WebSocket("ws://localhost:8081/ws/1")

    this.connection.onmessage = function(event) {
      console.log(event);
    }

    this.connection.onopen = function(event) {
      console.log(event)
      console.log("Successfully connected to the echo websocket server...")
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
