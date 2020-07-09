<template>
<div id="home">
  <Title msg="Not ScribbleIO"/>
  <CreateButton @onCreateRoom="onCreateRoom"/>
  <JoinModal @onJoinRoom="onJoinRoom"/>
  <Canvas/>
</div>
</template>

<script>
import Title from '../components/Title.vue'
import CreateButton from '../components/CreateButton.vue'
import JoinModal from '../components/JoinModal.vue'
import Canvas from '../components/Canvas.vue'

export default {
  name: 'Home',
  components: {
    Title,
    CreateButton,
    JoinModal,
    Canvas,
  },
  methods: {
    onCreateRoom: function(lobbyId) {
      this.$router.push({ name: 'lobby', params: { lobbyId: lobbyId }});
    },
    onJoinRoom: function(lobbyId) {
      this.$router.push({ name: 'lobby', params: { lobbyId: lobbyId }});
    }
  }
}
</script>