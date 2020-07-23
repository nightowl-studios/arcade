<template>
    <div class="player">
        <h2>{{ nickname }}</h2>
    </div>
</template>

<script>
export default {
    name: "Player",
    props: {
        nickname: String,
    }
};
</script>

<style scoped>
.player {
    border: solid 1px black;
    text-align: center;
    background-color: orange;
}
</style>
