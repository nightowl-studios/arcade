<template>
  <div id="app">
    <div v-if="isConnected">
      <img alt="Vue logo" src="./assets/logo.png">
      <HelloWorld msg="Welcome to Your Vue.js App"/>
      <b-button v-on:click="sendMessage('hello')">Button</b-button>
    </div>
    <div v-else>
      <h1>Lobby</h1>
      <ul>
        <li v-for="player in players" :key="player.name">
          {{ player.name }}
        </li>
      </ul>

    </div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld
  },
  data: function() {
    return {
      connection: null,
      isConnected: false,
      players: [
        { name: "Gordon" },
        { name: "Byron" },
        { name: "Zach" },
        { name: "Sam" }
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
