<template>
    <div class="player-list">
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
    background-color: lightblue;
    border: 1px solid black;
}
</style>
