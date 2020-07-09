<template>
<div id="home">
  <Title msg="Not ScribbleIO"/>
  <CreateButton @onCreateRoom="onCreateRoom"/>
  <JoinModal @onJoinRoom="onJoinRoom"/>
  <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
  <Canvas/>
</div>
</template>

<script>
import Title from '../components/Title.vue'
import CreateButton from '../components/CreateButton.vue'
import JoinModal from '../components/JoinModal.vue'
import Canvas from '../components/Canvas.vue'
import { EventBus } from '../eventBus.js';
import { ArcadeWebSocket } from '../webSocket.js';

export default {
  name: 'Home',
  components: {
    Title,
    CreateButton,
    JoinModal,
    Canvas,
  },
  data: function() {
    return {
      connectionState: "DISCONNECTED"
    }
  },
  methods: {
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
    EventBus.$on('connected', (lobbyId) => {
      console.log(lobbyId)
      this.connectionState = "CONNECTED";
      this.$router.push({ name: 'lobby', params: { lobbyId: lobbyId }}) 
    })
  }
}
</script>