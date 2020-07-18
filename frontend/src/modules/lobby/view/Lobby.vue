<template>
  <div id="lobby">
    <Header title="Welcome to Not Scribble" />
    <Header class="lobby-header-room-id" :title="lobbyId" />
    <div class="lobby-buttons">
      <b-button class="lobby-button" variant="success" v-on:click="goToScribble">Start game</b-button>
      <Nickname @onChangeNickname="onChangeNickname" />
      <b-button class="exit-button" variant="danger" v-on:click="exitToHome">Exit Lobby</b-button>
    </div>
    <PlayerList :players="players" />
  </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import Nickname from "../components/Nickname.vue";
//import { EventBus } from "@/eventBus.js";
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";

export default {
  mixins: [WebSocketMixin],
  name: "Lobby",
  components: {
    Header,
    PlayerList,
    Nickname
  },
  data: function() {
    return {
      lobbyId: "",
      connectionState: "DISCONNECTED"
    };
  },
  methods: {
    onChangeNickname: function(event) {
      let message = {
        api: "hub",
        payload: {
          changeNameTo: event.nickname
        }
      };
      this.$webSocketService.send(message);
    },
    sendPlayerMessage: function() {
      let message = {
        api: "hub",
        payload: {
          requestLobbyDetails: true
        }
      };
      this.$webSocketService.send(message);
    },
    goToScribble: function() {
      this.$router.push({ path: "/scribble/" + this.lobbyId });
    },
    exitToHome: function() {
      this.$webSocketService.disconnect();
      this.$router.push({ name: "home" });
    }
  }
};
</script>
<style scoped>
#lobby {
  padding-left: 100px;
  padding-right: 100px;
}

.lobby-button,
.exit-button {
  margin-left: 2px;
  margin-right: 2px;
}

.lobby-header-room-id {
  color: orange;
}
</style>
