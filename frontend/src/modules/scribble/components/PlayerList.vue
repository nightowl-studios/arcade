<template>
    <div id="player-list">
        <h2>Players</h2>
        <div v-for="player in players" :key="player">{{ player.nickname }}</div>
    </div>
</template>

<script>
export default {
    name: "PlayerList",
    props: {
        players: Array,
    },
};
</script>

<style scoped>
#player-list {
    flex: 1;
    background-color: lightblue;
}
</style>
