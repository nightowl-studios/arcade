<template>
<div>
    <LobbyText :clients="clients"/>
    <div>{{connectionState}}</div>
    <div>Room Id: {{ lobbyId }}</div>
    <Nickname @onChangeNickname="onChangeNickname"/>
    <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
</div>
</template>

<script>
import LobbyText from '../components/LobbyText.vue'
import Nickname from '../components/Nickname.vue'
import { EventBus } from '../eventBus.js';
import { ArcadeWebSocket } from '../webSocket.js';
import axios from 'axios';

export default {
  name: 'Lobby',
  components: {
    LobbyText,
    Nickname
  },
  data: function() {
    return {
      clients: [],
      lobbyId: "",
      connectionState: "DISCONNECTED"
    }
  },
  methods: {
    onChangeNickname: function(event) {
      let message = {
        "api":"hub",
        "payload":{
          "changeNameTo": event.nickname
        }
      }
      ArcadeWebSocket.send(message);
    },
    sendPlayerMessage: function() {
      let message = {
        "api":"hub",
        "payload":{
          "requestLobbyDetails":true
        }
      }
      ArcadeWebSocket.send(message);
    }
  },
  created() {
    this.lobbyId = this.$router.currentRoute.params.lobbyId;

    if (!ArcadeWebSocket.isConnected()) {
      let apiUrl = this.$httpURL + '/hub' + '/' + this.lobbyId;
      axios
        .get(apiUrl)
        .then(response => {
          if (response.data.exists) {
            ArcadeWebSocket.connect(this.lobbyId);
          } else {
            console.log("HubId does not exist...");
            // TODO display error here
            this.$router.push({ path: '/' })
          }
        });
    } else {
      this.connectionState = "CONNECTED";
    }
    
    EventBus.$on('connected', () => {
      this.connectionState = "CONNECTED";
    }),
    EventBus.$on(this.$hubAPI, (data) => {
      this.clients = data.connectedClients;
    }) 
  }
}
</script>