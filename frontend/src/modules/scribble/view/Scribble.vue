<template>
  <div id="scribble">
    <CanvasPanel :colors="colors" :sizes="sizes" />
    <Gameroom :clients="players" />
    <Nickname />
  </div>
</template>

<script>
import CanvasPanel from "../components/CanvasPanel.vue";
import Gameroom from "../components/Gameroom.vue";
import Nickname from "@/modules/lobby/components/Nickname.vue";
//import { EventBus } from "@/eventBus.js";
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";

export default {
  mixins: [WebSocketMixin],
  name: "Scribble",
  data: function() {
    return {
      colors: ["#000000", "#4287f5", "#da42f5", "#7ef542"],
      sizes: [8, 16, 32, 64]
    };
  },
  components: {
    CanvasPanel,
    Gameroom,
    Nickname
  }
};
</script>
