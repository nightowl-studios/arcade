<template>
  <div>
    <canvas ref="canvas" width="400" height="400"></canvas>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      mouseDown: false,
      previousPosition: { x: 0, y: 0 },
      lineStyle: { strokeStyle: "black", lineWidth: 2 },
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
        this.draw(this.previousPosition, currentPosition, this.lineStyle);
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

    draw: function(from, to, lineStyle) {
      this.context.beginPath();
      this.context.moveTo(from.x, from.y);
      this.context.lineTo(to.x, to.y);
      this.context.strokeStyle = lineStyle.strokeStyle;
      this.context.lineWidth = lineStyle.lineWidth;
      this.context.stroke();
      this.context.closePath();
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