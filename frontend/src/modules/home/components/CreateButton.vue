<template>
  <div>
    <b-button variant="success" v-on:click="createRoom">Create</b-button>
  </div>
</template>

<script>
export default {
  name: "CreateButton",
  methods: {
    createRoom: async function() {
      console.log("Creating room...");
      let lobbyId = await this.$hubApiService.createLobby();
      this.$emit("onCreateRoom", lobbyId);
    }
  }
};
</script>
