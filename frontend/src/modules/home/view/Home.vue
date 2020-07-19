<template>
  <div id="home">
    <Title id="title" msg="Not ScribbleIO" />
    <NicknameInput ref="nicknameInput" />
    <CreateButton id="createButton" @onCreateRoom="onCreateRoom" />
    <JoinRoom id="joinRoom" @onJoinRoom="onJoinRoom" />
  </div>
</template>

<script>
import Title from "../components/Title.vue";
import CreateButton from "../components/CreateButton.vue";
import JoinRoom from "../components/JoinRoom.vue";
import { EventBus } from "@/eventBus.js";
import NicknameInput from "@/modules/common/components/NicknameInput.vue";

export default {
  name: "Home",

  components: {
    Title,
    CreateButton,
    JoinRoom,
    NicknameInput
  },

  methods: {
    onCreateRoom: function(lobbyId) {
      this.connectToRoom(lobbyId);
    },

    onJoinRoom: function(lobbyId) {
      this.connectToRoom(lobbyId);
    },

    connectToRoom: function(lobbyId) {
      if (!this.$refs["nicknameInput"].validateNickname()) {
        return;
      }
      this.$webSocketService.connect(lobbyId);
    }
  },
  created() {
    EventBus.$on("connected", lobbyId => {
      this.$router.push({ name: "lobby", params: { lobbyId: lobbyId } });
      this.$refs["nicknameInput"].changeNickname();
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
