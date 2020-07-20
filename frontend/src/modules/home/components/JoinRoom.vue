<template>
  <div id="joinRoom">
    <div id="joinRoomInput">
      <b-form-input id="lobbyIdInput" v-model="lobbyId" maxlength="4" placeholder="Lobby ID"></b-form-input>
      <b-button id="joinButton" variant="success" v-on:click="onJoin">Join</b-button>
    </div>
    <p
      id="hubIdDoesNotExistError"
      :style="{visibility: showError ? 'visible' : 'hidden'}"
    >Lobby doesn't exist!</p>
  </div>
</template>

<script>
export default {
  name: "joinRoom",
  data: function() {
    return {
      lobbyId: "",
      showError: false
    };
  },
  watch: {
    lobbyId: function(val) {
      this.lobbyId = val.toUpperCase();
      this.showError = false;
    }
  },
  methods: {
    onJoin: async function() {
      if (this.lobbyId.length !== 4) {
        this.showError = true;
        return;
      }
      const lobbyExists = await this.$hubApiService.checkLobbyExists(
        this.lobbyId
      );
      if (lobbyExists) {
        this.$emit("onJoinRoom", this.lobbyId);
      } else {
        this.showError = true;
      }
    }
  }
};
</script>

<style scoped>
#joinRoomInput {
  display: flex;
}

#lobbyIdInput {
  margin-right: 10px;
  width: 90px;
}

#hubIdDoesNotExistError {
  display: inline-block;
  margin-top: 8px;
  color: red;
}
</style>
