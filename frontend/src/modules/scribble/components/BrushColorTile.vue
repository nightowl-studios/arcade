<template>
    <div :style="style" v-on:click="onTileClicked"></div>
</template>

<script>
export default {
    name: "BrushColorTile",
    props: {
        color: String,
    },
    computed: {
        style() {
            return {
                height: "20px",
                width: "20px",
                borderRadius: "30%",
                backgroundColor: this.color,
                float: "left",
            };
        },
    },
    methods: {
        onTileClicked() {
            this.$emit("colorSelected", this.color);
        },
    },
};
</script>

<style scoped></style>
