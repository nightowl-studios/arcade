<template>
<div id="lobby">
    <h1>Lobby</h1>
    <div v-for="player in clients" :key="player">
        <Player :id="player.clientUUID.UUID" :name="player.nickname"/>
    </div>
</div>
</template>

<script>
import Player from './Player.vue'

export default {
  name: 'Lobby',
  props: {
    clients: Array,
  },
  components: {
    Player,
  },
}
</script>