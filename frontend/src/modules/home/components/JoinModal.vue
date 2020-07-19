<template>
  <div id="joinModal">
    <b-button variant="success" v-b-modal.modal-1>Join</b-button>
    <b-modal id="modal-1" @ok="onOKClicked" title="BootstrapVue">
      <input v-model="lobbyId" placeholder="Enter lobby id" />
    </b-modal>
  </div>
</template>

<script>
export default {
  name: "JoinModal",
  data: function() {
    return {
      lobbyId: ""
    };
  },
  methods: {
    onOKClicked: async function() {
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
