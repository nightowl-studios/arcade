<template>
    <div class="canvas-panel">
        <Canvas
            ref="canvas"
            :width="2000"
            :height="2000"
            :defaultBrushStyle="defaultBrushStyle"
            :drawingLocked="isCanvasLocked"
            @drawAction="sendDrawAction"
            @requestHistory="sendRequestHistory"
        />
        <BrushSelector :colors="colors" :sizes="sizes" />
    </div>
</template>

<script>
import Canvas from "./Canvas.vue";
import BrushSelector from "./BrushSelector.vue";
import { createBrushStyle } from "../utility/BrushStyleUtils";
import { Event } from "@/events.js";
import { EventBus } from "@/eventBus.js";

export default {
    name: "CanvasPanel",

    props: {
        colors: Array,
        sizes: Array,
        isCanvasLocked: Boolean,
    },

    components: {
        Canvas,
        BrushSelector,
    },

    computed: {
        defaultBrushStyle() {
            return createBrushStyle(this.sizes[0], this.colors[0]);
        },
    },

    mounted: function () {
        EventBus.$on(Event.CANVAS_UPDATE, this.handleDrawMessage);
        this.sendRequestHistory();
    },

    methods: {
        sendDrawAction(drawAction) {
            this.$scribbleGameController.draw(drawAction);
        },

        sendRequestHistory() {
            this.$scribbleGameController.requestDrawHistory();
        },

        handleDrawMessage(drawMessage) {
            if (drawMessage.history != null) {
                for (const action of drawMessage.history) {
                    this.$refs["canvas"].draw(action);
                }
            }
            if (drawMessage.action != null) {
                this.$refs["canvas"].draw(drawMessage.action);
            }
        },
    },
};
</script>
