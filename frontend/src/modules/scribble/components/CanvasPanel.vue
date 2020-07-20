<template>
    <div>
        <Canvas
            ref="canvas"
            :width="400"
            :height="400"
            :defaultBrushStyle="defaultBrushStyle"
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
import { EventBus } from "@/eventBus.js";
import { createDrawMessage } from "@/modules/common/utility/WebSocketMessageUtils";

export default {
    name: "CanvasPanel",

    props: {
        colors: Array,
        sizes: Array,
    },

    components: {
        Canvas,
        BrushSelector,
    },

    methods: {
        sendDrawAction(drawAction) {
            const drawMsg = createDrawMessage({
                action: drawAction,
                requestHistory: false,
            });
            this.$webSocketService.send(drawMsg);
        },

        sendRequestHistory() {
            const requestHistoryMsg = createDrawMessage({
                requestHistory: true,
            });
            this.$webSocketService.send(requestHistoryMsg);
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

<style scoped></style>
