<template>
    <div>
        <canvas ref="canvas" :width="width" :height="height"></canvas>
    </div>
</template>

<script>
import { EventBus } from "@/eventBus.js";

export default {
    name: "Canvas",

    props: {
        width: Number,
        height: Number,
        defaultBrushStyle: Object,
        drawingLocked: Boolean,
    },

    data: function () {
        return {
            mouseDown: false,
            previousPosition: { x: 0, y: 0 },
            brushStyle: this.defaultBrushStyle,
            canvas: null,
            context: null,
            offsetLeft: 0,
            offsetTop: 0,
        };
    },

    mounted: function () {
        this.canvas = this.$refs["canvas"];
        this.context = this.canvas.getContext("2d");
        this.updateOffset();
        this.canvas.addEventListener("mousemove", this.onMouseMove, false);
        this.canvas.addEventListener("mousedown", this.onMouseDown, false);
        this.canvas.addEventListener("mouseup", this.onMouseUp, false);
        this.canvas.addEventListener("mouseover", this.onMouseOver, false);
        window.addEventListener("scroll", this.updateOffset, false);
        window.addEventListener("resize", this.updateOffset, false);
        EventBus.$on("brushUpdated", this.setBrushStyle);
    },

    methods: {
        onMouseDown: function (event) {
            this.previousPosition = {
                x: event.clientX - this.offsetLeft,
                y: event.clientY - this.offsetTop,
            };
            this.handleDrawInput(
                this.previousPosition,
                this.previousPosition,
                this.brushStyle
            );
            this.mouseDown = true;
        },

        onMouseMove: function (event) {
            if (this.mouseDown) {
                let currentPosition = {
                    x: event.clientX - this.offsetLeft,
                    y: event.clientY - this.offsetTop,
                };
                this.handleDrawInput(
                    this.previousPosition,
                    currentPosition,
                    this.brushStyle
                );
                this.previousPosition = currentPosition;
            }
        },

        onMouseUp: function () {
            this.mouseDown = false;
        },

        onMouseOver: function (event) {
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
                    y: event.clientY - this.canvas.offsetTop,
                };
            }
        },

        updateOffset: function () {
            var br = this.canvas.getBoundingClientRect();
            this.offsetLeft = br.left;
            this.offsetTop = br.top;
        },

        setBrushStyle: function (brushStyle) {
            this.brushStyle = brushStyle;
        },

        handleDrawInput: function (from, to, brushStyle) {
            if (!this.drawingLocked) {
                let drawAction = {
                    from: from,
                    to: to,
                    brushStyle: brushStyle,
                    lineCap: this.context.lineCap,
                };
                this.draw(drawAction);
                this.$emit("drawAction", drawAction);
            }
        },

        draw: function (drawAction) {
            this.context.beginPath();
            this.context.moveTo(drawAction.from.x, drawAction.from.y);
            this.context.lineTo(drawAction.to.x, drawAction.to.y);
            this.context.strokeStyle = drawAction.brushStyle.brushColor;
            this.context.lineWidth = drawAction.brushStyle.brushSize;
            this.context.lineCap = "round";
            this.context.stroke();
            this.context.closePath();
        },
    },
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
