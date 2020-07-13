<template>
  <div id="home">
    <Title msg="Not ScribbleIO" />
    <CreateButton @onCreateRoom="onCreateRoom" />
    <JoinModal @onJoinRoom="onJoinRoom" />
    <div>{{ message }}</div>
  </div>
</template>

<script>
import Title from "../components/Title.vue";
import CreateButton from "../components/CreateButton.vue";
import JoinModal from "../components/JoinModal.vue";
import { ArcadeWebSocket } from "../webSocket.js";
import { EventBus } from "../eventBus.js";
import { mapState } from "vuex";

export default {
  name: "Home",
  components: {
    Title,
    CreateButton,
    JoinModal
  },
  computed: {
    ...mapState({
      message: state => state.application.message
    })
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
    EventBus.$on("connected", lobbyId => {
      this.$router.push({ name: "lobby", params: { lobbyId: lobbyId } });
    });
  }
};
</script>
