<template>
    <div class="word">{{ action }} {{ word }}</div>
</template>

<script>
export default {
    name: "Header",
    props: {
        word: String,
        isGuessing: Boolean,
    },
    computed: {
        action() {
            return this.isGuessing ? "Guessing" : "Drawing";
        },
    },
};
</script>

<style lang="scss" scoped>
.word {
    text-align: center;
}
</style>
