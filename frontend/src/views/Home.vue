<template>
<div id="home">
  <Title msg="Not ScribbleIO"/>
  <CreateButton @onCreateRoom="onCreateRoom"/>
  <JoinModal @onJoinRoom="onJoinRoom"/>
</div>
</template>

<script>
import Title from '../components/Title.vue'
import CreateButton from '../components/CreateButton.vue'
import JoinModal from '../components/JoinModal.vue'
import { ArcadeWebSocket } from '../webSocket.js'
import { EventBus } from '../eventBus.js'

export default {
  name: 'Home',
  components: {
    Title,
    CreateButton,
    JoinModal
  },
  methods: {
    onCreateRoom: function(lobbyId) {
      ArcadeWebSocket.connect(lobbyId);
    },
    onJoinRoom: function(lobbyId) {
      ArcadeWebSocket.connect(lobbyId);
    }
  },
  created() {
    EventBus.$on('connected', (lobbyId) => {
      this.$router.push({ name: 'lobby', params: { lobbyId: lobbyId }});
    });
  }
}
</script>