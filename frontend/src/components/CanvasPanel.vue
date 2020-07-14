<template>
  <div>
    <Canvas :width="400" :height="400" :defaultBrushStyle="defaultBrushStyle" />
    <BrushSelector :colors="colors" :sizes="sizes" />
  </div>
</template>

<script>
import Canvas from "./Canvas.vue";
import BrushSelector from "./BrushSelector.vue";
import { createBrushStyle } from "../utility/BrushStyleUtils";
import { createBrushStrokeMessage } from "../utility/WebSocketMessageUtils";
import { EventBus } from "../eventBus.js";

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

  mounted: function() {
    EventBus.$on("brushStroke", this.onBrushStroke);
  },

  computed: {
    defaultBrushStyle() {
      return createBrushStyle(this.sizes[0], this.colors[0]);
    }
  },

  methods: {
    onBrushStroke: function(brushStroke) {
      let message = createBrushStrokeMessage(brushStroke);
      this.$webSocketService.send(message);
    }
  }
};
</script>

<style scoped>
</style>
