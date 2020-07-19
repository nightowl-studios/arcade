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
    <b-container fluid class="invitation-link-section">
      <b-row align-v="center" class="justify-content-md-center">
        <b-col md="auto">Invitation Link: </b-col>
        <b-col md="auto"><input id="invitation-link" type='text' v-model="inviteLink" readonly></b-col>
        <b-col md="auto"><b-button class="copy-link-button" variant="success" v-on:click="copyInviteLink">Copy Link</b-button></b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import Nickname from "../components/Nickname.vue";
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
      inviteLink: window.location.href
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
      let message = {
        api: "game",
        payload: {
          gameMasterAPI: "waitForStart",
          waitForStart: {
            startGame: true
          }
        }
      };
      this.$webSocketService.send(message);
      this.$router.push({ path: "/scribble/"  + this.lobbyId });
    },
    exitToHome: function() {
      this.$webSocketService.disconnect();
      this.$router.push({ name: "home" });
    },
    copyInviteLink: function() {
      let linkToCopy = document.querySelector('#invitation-link');
      linkToCopy.select();
      try {
        var successful = document.execCommand('copy');
        var msg = successful ? 'successful' : 'unsuccessful';
        alert('Invitation link was copied ' + msg);
      } catch (err) {
        alert('Oops, unable to copy');
      }
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

.invitation-link-section {
  margin-top: 25px;
}

#invitation-link {
  color: black;
  width: 300px;
}
</style>
