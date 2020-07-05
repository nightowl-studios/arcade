<template>
<div id="lobby">
    <h1>Lobby</h1>
    <div v-for="player in players" :key="player">
        <Player :name="player.name" :id="player.id" />
    </div>
</div>
</template>

<script>
import Player from './Player.vue'

export default {
  name: 'Lobby',
  components: {
    Player
  },
  data: function() {
    return {
      players: [
        { name: "Gordon", id: "ID12345"},
        { name: "Byron", id: "ID12346"},
        { name: "Zach", id: "ID12347" },
        { name: "Sam", id: "ID12348" }
      ],
    }
  }
}
</script>