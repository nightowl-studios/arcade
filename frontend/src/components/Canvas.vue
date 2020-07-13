<template>
  <div>
    <canvas ref="canvas" :width="width" :height="height"></canvas>
  </div>
</template>

<script>
import { EventBus } from "../eventBus.js";

export default {
  name: "Canvas",

  props: {
    width: Number,
    height: Number,
    defaultBrushStyle: Object
  },

  data: function() {
    return {
      mouseDown: false,
      previousPosition: { x: 0, y: 0 },
      brushStyle: this.defaultBrushStyle,
      canvas: null,
      context: null
    };
  },

  mounted: function() {
    this.canvas = this.$refs["canvas"];
    this.context = this.canvas.getContext("2d");
    this.canvas.addEventListener("mousemove", this.onMouseMove, false);
    this.canvas.addEventListener("mousedown", this.onMouseDown, false);
    this.canvas.addEventListener("mouseup", this.onMouseUp, false);
    this.canvas.addEventListener("mouseover", this.onMouseOver, false);
    EventBus.$on("brushUpdated", this.setBrushStyle);
  },

  methods: {
    onMouseDown: function(event) {
      this.previousPosition = {
        x: event.clientX - this.canvas.offsetLeft,
        y: event.clientY - this.canvas.offsetTop
      };
      this.mouseDown = true;
    },

    onMouseMove: function(event) {
      if (this.mouseDown) {
        let currentPosition = {
          x: event.clientX - this.canvas.offsetLeft,
          y: event.clientY - this.canvas.offsetTop
        };
        this.draw(this.previousPosition, currentPosition, this.brushStyle);
        this.previousPosition = currentPosition;
      }
    },

    onMouseUp: function() {
      this.mouseDown = false;
    },

    onMouseOver: function(event) {
      if (
        !(event.buttons === undefined
          ? (event.which & 1) === 1
          : (event.buttons & 1) === 1)
      ) {
        this.mouseDown = false;
      } else {
        this.mouseDown = true;
        this.previousPosition = {
          x: event.clientX - this.canvas.offsetLeft,
          y: event.clientY - this.canvas.offsetTop
        };
      }
    },

    setBrushStyle: function(brushStyle) {
      this.brushStyle = brushStyle;
    },

    draw: function(from, to, brushStyle) {
      this.context.beginPath();
      this.context.moveTo(from.x, from.y);
      this.context.lineTo(to.x, to.y);
      this.context.strokeStyle = brushStyle.brushColor;
      this.context.lineWidth = brushStyle.brushSize;
      this.context.lineCap = "round";
      this.context.stroke();
      this.context.closePath();

      EventBus.$emit("brushStroke", {
        from: from,
        to: to,
        brushStyle: brushStyle,
        lineCap: this.context.lineCap
      });
    }
  }
};
</script>

<style scoped>
canvas {
  position: absolute;
  top: 10%;
  left: 10%;
  border: 2px solid;
}
</style>
