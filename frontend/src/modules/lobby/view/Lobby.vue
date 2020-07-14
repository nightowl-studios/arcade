<template>
  <div id="lobby">
    <Header title="Welcome to Not Scribble" />
    <div class="lobby-buttons">
      <b-button class="lobby-button" variant="success">Start game</b-button>
      <b-button class="lobby-button" variant="primary">Change nickname</b-button>
    </div>
    <Header :title="lobbyId" />
    <PlayerList :players="players" />
    <div>{{connectionState}}</div>
    <div>Room Id: {{ lobbyId }}</div>
    <Nickname @onChangeNickname="onChangeNickname" />
    <b-button v-on:click="sendPlayerMessage()">Send a Message</b-button>
    <b-button v-on:click="goToScribble()">Go to Scribble</b-button>
  </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import Nickname from "../components/Nickname.vue";
//import { EventBus } from "@/eventBus.js";

export default {
  name: "Lobby",
  components: {
    Header,
    PlayerList,
    Nickname
  },
  data: function() {
    return {
      players: [
        {
          nickname: "Gordon"
        },
        {
          nickname: "Sam"
        },
        {
          nickname: "Byron"
        }
      ],
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
      this.$router.push({ path: "/scribble" });
    }
  },
  async created() {
    // EventBus.$on("connected", () => {
    //   this.connectionState = "CONNECTED";
    // }),
    //   EventBus.$on(this.$hubAPI, data => {
    //     this.players = data.connectedClients;
    //   });
    // this.lobbyId = this.$router.currentRoute.params.lobbyId;
    // if (!this.$webSocketService.isConnected()) {
    //   let lobbyExists = await this.$hubApiService.checkLobbyExists(
    //     this.lobbyId
    //   );
    //   if (!lobbyExists) {
    //     this.$router.push({ name: "404" });
    //   } else {
    //     this.$webSocketService.connect(this.lobbyId);
    //   }
    // }
  }
};
</script>
<style scoped>
#lobby {
  padding-left: 100px;
  padding-right: 100px;
}

.lobby-button {
  margin-left: 2px;
  margin-right: 2px;
}
</style>
