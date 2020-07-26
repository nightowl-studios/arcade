<template>
    <div class="box" v-on:click="onTileClicked">
        <div :style="style"></div>
    </div>
</template>

<script>
export default {
    name: "BrushSizeTile",
    props: {
        size: Number,
    },
    computed: {
        style() {
            return {
                height: `${this.size / 2}px`,
                width: `${this.size / 2}px`,
                position: "absolute",
                backgroundColor: "black",
                borderRadius: "50%",
                margin: `${(40 - this.size / 2) / 2 - 1}px`,
            };
        },
    },
    methods: {
        onTileClicked() {
            this.$emit("sizeSelected", this.size);
        },
    },
};
</script>

<style lang="scss" scoped>
.box {
    height: 40px;
    width: 40px;
    text-align: center;
    border: 1px solid;
}
</style>
