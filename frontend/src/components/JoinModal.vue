<template>
  <div id="joinModal">
    <b-button variant="success" v-b-modal.modal-1>Join</b-button>
    <b-modal id="modal-1" @ok="onOKClicked" title="BootstrapVue">
      <input v-model="lobbyId" placeholder="Enter lobby id">
    </b-modal>
  </div>
</template>

<script>
import axios from 'axios';
import { ArcadeWebSocket } from '../webSocket.js'

export default {
  name: "JoinModal",
  data: function() {
    return {
      lobbyId: ''
    }
  },
  methods: {
    onOKClicked: function() {
      console.log("Joining room " + this.lobbyId + "...")
      let apiUrl = this.$httpURL + '/hub' + '/' + this.lobbyId;
      axios
        .get(apiUrl)
        .then(response => {
          if (response.data.exists) {
            ArcadeWebSocket.connect(this.lobbyId);
            this.$emit('onJoinRoom', this.lobbyId);
          } else {
            console.log("HubId does not exist...");
          }
        });
    }
  }
}
</script>

<style scoped>
#joinModal {
  margin-top: 15px;
}
</style>
