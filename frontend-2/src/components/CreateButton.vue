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
      // TODO make api request to backend
      let apiUrl = "http://localhost:8080/hub";
      axios
        .get(apiUrl)
        .then(response => (this.$emit("onCreateRoom", response)));
    }
  }
}
</script>