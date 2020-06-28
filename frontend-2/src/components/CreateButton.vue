<template>
  <div>
    <b-button v-on:click="createRoom">Create</b-button>
  </div>
</template>

<script>
export default {
  name: "CreateButton",
  methods: {
    createRoom: function() {
      console.log("Creating room...")
      // TODO make api request to backend
      let apiUrl = "";
      axios
        .get(apiUrl)
        .then(response => (this.$emit("onCreateRoom", response)));
    }
  }
}
</script>