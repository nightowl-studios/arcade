<template>
    <div class="round-number">
        <span> Round {{ round }} </span>
    </div>
</template>

<script>
export default {
    name: "RoundNumber",
    computed: {
        round() {
            return this.$store.getters["scribble/getRoundNumber"];
        },
    },
};
</script>

<style lang="scss" scoped>
.round-number {
    background-color: #e5f4fc;
    border: 1px solid #a1c5d8;
    padding-top: 5px;
    padding-bottom: 5px;
    border-radius: 6px;
    margin: 5px auto 5px auto;
    width: 90px;
}
</style>
