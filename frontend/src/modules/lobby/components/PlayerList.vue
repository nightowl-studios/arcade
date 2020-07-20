<template>
    <div class="player-list">
        <h1>Players</h1>
        <div v-for="player in players" :key="player">
            <Player :name="player.nickname" />
        </div>
    </div>
</template>

<script>
import Player from './Player.vue'

export default {
    name: 'PlayerList',
    props: {
        players: Array,
    },
    components: {
        Player,
    },
}
</script>

<style scoped>
.player-list {
    background-color: orangered;
    margin-top: 10px;
}
</style>
