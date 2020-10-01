<template>
    <div class="brush-selector">
        <div class="brush-selector__color"></div>
        <BrushColorTile
            v-for="color in colors"
            :key="color"
            :color="color"
            @colorSelected="onColorSelected"
        />
        <div class="brush-selector__size">
            <BrushSizeTile
                v-for="size in sizes"
                :key="size"
                :size="size"
                @sizeSelected="onSizeSelected"
            />
        </div>
        <div class="brush-selector__actions">
            <BrushResetActionTile
                :v-if="showResetButton"
                @resetCanvas="onResetCanvas"
            />
        </div>
    </div>
</template>

<script>
import BrushSizeTile from "./BrushSizeTile.vue";
import BrushColorTile from "./BrushColorTile.vue";
import BrushResetActionTile from "./BrushResetActionTile.vue";
import { EventBus } from "@/eventBus.js";
import { createBrushStyle } from "../utility/BrushStyleUtils";
import { DrawEvent } from "@/modules/scribble/utility/drawEvents";

export default {
    name: "BrushSelector",

    components: {
        BrushSizeTile,
        BrushColorTile,
        BrushResetActionTile,
    },

    props: {
        colors: Array,
        sizes: Array,
        showResetButton: Boolean,
    },

    data: function () {
        return {
            currentSize: this.sizes[0],
            currentColor: this.colors[0],
        };
    },

    methods: {
        onSizeSelected: function (size) {
            this.currentSize = size;
            this.emitUpdatedBrush();
        },
        onColorSelected: function (color) {
            this.currentColor = color;
            this.emitUpdatedBrush();
        },
        onResetCanvas: function() {
            EventBus.$emit(DrawEvent.RESET_CANVAS);
        },
        emitUpdatedBrush: function () {
            EventBus.$emit(
                DrawEvent.UPDATE_BRUSH,
                createBrushStyle(this.currentSize, this.currentColor)
            );
        },
    },
};
</script>

<style lang="scss" scoped>
.brush-selector {
    display: flex;

    &__size {
        display: flex;
    }
}
</style>
