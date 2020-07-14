<template>
  <div id="scribble">
    <CanvasPanel :colors="colors" :sizes="sizes" />
    <Gameroom :clients="clients" />
  </div>
</template>

<script>
import CanvasPanel from "../components/CanvasPanel.vue";
import Gameroom from "../components/Gameroom.vue";
import { EventBus } from "@/eventBus.js";

export default {
  name: "Scribble",
  data: function() {
    return {
      clients: [],
      colors: ["#000000", "#4287f5", "#da42f5", "#7ef542"],
      sizes: [8, 16, 32, 64]
    };
  },
  components: {
    CanvasPanel,
    Gameroom
  },
  created() {
    EventBus.$on(this.$hubAPI, data => {
      this.clients = data.connectedClients;
    });

    let message = {
      api: "hub",
      payload: {
        requestLobbyDetails: true
      }
    };
    this.$webSocketService.send(message);
  }
};
</script>
