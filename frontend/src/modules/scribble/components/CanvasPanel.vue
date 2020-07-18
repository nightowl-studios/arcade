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
};
</script>

<style scoped>
</style>
