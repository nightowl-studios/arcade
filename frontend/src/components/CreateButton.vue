<template>
  <div>
    <b-button variant="success" v-on:click="createRoom">Create</b-button>
  </div>
</template>

<script>
import axios from 'axios';

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
          this.$emit("onCreateRoom", lobbyId);
        });
    }
  }
}
</script>