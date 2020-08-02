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
                border: "1px #000000 solid",
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
