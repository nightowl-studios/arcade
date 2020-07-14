<template>
  <div id="brush-selector">
    <div id="color-selector"></div>
    <BrushColorTile
      v-for="color in colors"
      :key="color"
      :color="color"
      @colorSelected="onColorSelected"
    />
    <div id="size-selector">
      <BrushSizeTile v-for="size in sizes" :key="size" :size="size" @sizeSelected="onSizeSelected" />
    </div>
  </div>
</template>

<script>
import BrushSizeTile from "./BrushSizeTile.vue";
import BrushStyleFactory from "../mixins/BrushStyleFactory";
import BrushColorTile from "./BrushColorTile.vue";
import { EventBus } from "../eventBus.js";

export default {
  name: "BrushSelector",

  components: {
    BrushSizeTile,
    BrushColorTile
  },

  props: {
    colors: Array,
    sizes: Array
  },

  data: function() {
    return {
      currentSize: this.sizes[0],
      currentColor: this.colors[0]
    };
  },

  mixins: [BrushStyleFactory],

  methods: {
    onSizeSelected: function(size) {
      this.currentSize = size;
      this.emitUpdatedBrush();
    },
    onColorSelected: function(color) {
      this.currentColor = color;
      this.emitUpdatedBrush();
    },
    emitUpdatedBrush: function() {
      EventBus.$emit(
        "brushUpdated",
        this.createBrushStyle(this.currentSize, this.currentColor)
      );
    }
  }
};
</script>

<style scoped>
#brush-selector {
  width: 500px;
  height: 500px;
}
#size-selector {
  width: 500px;
  height: 500px;
}
</style>
