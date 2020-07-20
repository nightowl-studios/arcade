<template>
    <div :style="style" v-on:click="onTileClicked"></div>
</template>

<script>
export default {
    name: 'BrushSizeTile',
    props: {
        size: Number,
    },
    computed: {
        style() {
            return {
                height: `${this.size}px`,
                width: `${this.size}px`,
                borderRadius: '50%',
                backgroundColor: 'black',
            }
        },
    },
    methods: {
        onTileClicked() {
            this.$emit('sizeSelected', this.size)
        },
    },
}
</script>

<style scoped></style>
