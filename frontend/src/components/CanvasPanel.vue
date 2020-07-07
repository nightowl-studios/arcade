<template>
  <div>
    <Canvas ref="canvas" />
    <BrushSelector :colors="colors" :sizes="sizes" @brushUpdated="onBrushUpdated" />
  </div>
</template>

<script>
import Canvas from "./Canvas.vue";
import BrushSelector from "./BrushSelector.vue";

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
  methods: {
    onBrushUpdated: function(brushStyle) {
      this.$refs["canvas"].setBrushStyle(brushStyle);
    }
  }
};
</script>

<style scoped>
</style>