<template>
  <div>
    <Canvas
      ref="canvas"
      :width="400"
      :height="400"
      :defaultBrushStyle="defaultBrushStyle"
      @drawAction="sendDrawAction"
    />
    <BrushSelector :colors="colors" :sizes="sizes" />
  </div>
</template>

<script>
import Canvas from "./Canvas.vue";
import BrushSelector from "./BrushSelector.vue";
import { createBrushStyle } from "../utility/BrushStyleUtils";
import { EventBus } from "@/eventBus.js";
import { createDrawActionMessage } from "@/modules/common/utility/WebSocketMessageUtils";

export default {
  name: "CanvasPanel",

  props: {
    colors: Array,
    sizes: Array
  },

  components: {
    Canvas,
    BrushSelector
  },

  computed: {
    defaultBrushStyle() {
      return createBrushStyle(this.sizes[0], this.colors[0]);
    }
  },

  mounted: function() {
    EventBus.$on("draw", this.handleDrawMessage);
  },

  methods: {
    sendDrawAction(drawAction) {
      this.$webSocketService.send(createDrawActionMessage(drawAction));
    },

    handleDrawMessage(drawMessage) {
      this.$ref["canvas"].draw(drawMessage.action);
    }
  }
};
</script>

<style scoped>
</style>
