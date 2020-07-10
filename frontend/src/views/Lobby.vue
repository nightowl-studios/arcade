<template>
<div>
    <LobbyText :clients="clients"/>
    <div>{{connectionState}}</div>
    <div>Room Id: {{ lobbyId }}</div>
    <Nickname @onChangeNickname="onChangeNickname"/>
    <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
    <Gameroom :clients="clients"/>
</div>
</template>

<script>
import LobbyText from '../components/LobbyText.vue'
import Nickname from '../components/Nickname.vue'
import Gameroom from '../components/Gameroom.vue'
import { EventBus } from '../eventBus.js';
import { ArcadeWebSocket } from '../webSocket.js';
import axios from 'axios';

export default {
  name: 'Lobby',
  components: {
    LobbyText,
    Nickname,
    Gameroom
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
    EventBus.$on('connected', () => {
      this.connectionState = "CONNECTED";
    }),
    EventBus.$on(this.$hubAPI, (data) => {
      this.clients = data.connectedClients;
    })

    this.lobbyId = this.$router.currentRoute.params.lobbyId;
    let lobbyExistsApiUrl = this.$httpURL + '/hub' + '/' + this.lobbyId;
    if (!ArcadeWebSocket.isConnected()) {
      axios
      .get(lobbyExistsApiUrl)
      .then(response => {
        if (!response.data.exists) {
          this.$router.push({ name: "404" });
        } else {
          ArcadeWebSocket.connect(this.lobbyId);
        }
      });
    }
  }
}
</script>