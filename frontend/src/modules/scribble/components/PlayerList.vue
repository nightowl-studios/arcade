<template>
    <div class="player-list">
        <h2>Players</h2>
        <div v-for="player in players" :key="player.uuid">
            <Player :player="player" />
        </div>
    </div>
</template>

<script>
import Player from "../components/Player.vue";

export default {
    name: "PlayerList",
    components: {
        Player,
    },
    props: {
        players: Array,
    },
};
</script>

<style scoped>
.player-list {
    flex: 1;
    background-color: lightblue;
}
</style>
