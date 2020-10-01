<template>
    <div v-on:click="onActionClicked">
        <b-icon icon="arrow-clockwise" font-scale="3"></b-icon>
    </div>
</template>

<script>
export default {
    name: "BrushResetActionTile",
    methods: {
        onActionClicked() {
            this.$emit("resetCanvas");
        },
    },
};
</script>
