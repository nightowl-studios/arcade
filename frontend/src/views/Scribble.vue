<template>
  <div id="scribble">
    <Canvas/>
    <Gameroom :clients="clients"/>
  </div>
</template>

<script>
import Canvas from '../components/Canvas.vue'
import Gameroom from '../components/Gameroom.vue'
import { EventBus } from '../eventBus.js';
import { ArcadeWebSocket } from '../webSocket.js';

export default {
  name: 'Scribble',
  data: function() {
    return {
      clients: []
    }
  },
  components: {
    Canvas,
    Gameroom
  },
  created() {
    EventBus.$on(this.$hubAPI, (data) => {
      this.clients = data.connectedClients;
    })

    let message = {
      "api":"hub",
      "payload":{
        "requestLobbyDetails":true
      }
    }
    ArcadeWebSocket.send(message);
  }
}
</script>