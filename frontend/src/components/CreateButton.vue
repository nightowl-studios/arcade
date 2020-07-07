<template>
  <div>
    <b-button variant="success" v-on:click="createRoom">Create</b-button>
  </div>
</template>

<script>
import axios from 'axios';
import { ArcadeWebSocket } from '../webSocket.js'

export default {
  name: "CreateButton",
  methods: {
    createRoom: function() {
      console.log("Creating room...")
      let apiUrl = this.$httpURL + '/hub';
      axios
        .get(apiUrl)
        .then(response => {
          let lobbyId = response.data.hubID;
          ArcadeWebSocket.connect(lobbyId);
          this.$emit("onCreateRoom", lobbyId);
        });
    }
  }
}
</script>