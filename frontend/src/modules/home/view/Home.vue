<template>
  <div id="home">
    <Title id="title" msg="Not ScribbleIO" />
    <CreateButton id="createButton" @onCreateRoom="onCreateRoom" />
    <JoinRoom id="joinRoom" @onJoinRoom="onJoinRoom" />
  </div>
</template>

<script>
import Title from "../components/Title.vue";
import CreateButton from "../components/CreateButton.vue";
import JoinRoom from "../components/JoinRoom.vue";
import { EventBus } from "@/eventBus.js";

export default {
  name: "Home",
  components: {
    Title,
    CreateButton,
    JoinRoom
  },
  methods: {
    onCreateRoom: function(lobbyId) {
      this.$webSocketService.connect(lobbyId);
    },
    onJoinRoom: function(lobbyId) {
      this.$webSocketService.connect(lobbyId);
    }
  },
  created() {
    EventBus.$on("connected", lobbyId => {
      this.$router.push({ name: "lobby", params: { lobbyId: lobbyId } });
    });
  }
};
</script>

<style scoped>
#home {
  display: grid;
  grid-template-rows: auto;
  justify-items: center;
}

#createButton {
  margin: 10px;
}

#joinRoom {
  margin: 10px;
}

</style>
