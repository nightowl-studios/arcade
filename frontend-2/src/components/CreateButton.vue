<template>
  <div>
    <b-button v-on:click="createRoom">Create</b-button>
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
        .then(response => (this.$emit("onCreateRoom", response)));
    }
  }
}
</script>