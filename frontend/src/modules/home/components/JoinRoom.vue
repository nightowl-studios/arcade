<template>
  <div id="joinRoom">
    <b-form-input id="lobbyIdInput" v-model="lobbyId" placeholder="Lobby ID"></b-form-input>
    <b-button id="joinButton" variant="success" v-on:click="onJoin">Join</b-button>
  </div>
</template>

<script>
export default {
  name: "joinRoom",
  data: function() {
    return {
      lobbyId: ""
    };
  },
  methods: {
    onJoin: async function() {
      console.log("Joining room " + this.lobbyId + "...");
      let lobbyExists = await this.$hubApiService.checkLobbyExists(
        this.lobbyId
      );
      if (lobbyExists) {
        this.$emit("onJoinRoom", this.lobbyId);
      } else {
        console.log("HubId does not exist...");
      }
    }
  }
};
</script>

<style scoped>
#joinRoom {
  display: flex;
}

#lobbyIdInput {
  margin-right: 10px;
  width: 90px;
}

</style>
