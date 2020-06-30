<template>
  <div>
    <b-button v-b-modal.modal-1>Launch demo modal</b-button>

    <b-modal id="modal-1" @ok="testMethod" title="BootstrapVue">
      <p class="my-4">Hello from modal!</p>
      <input v-model="roomId" placeholder="Enter room id">
    </b-modal>
  </div>
</template>

<script>
export default {
  name: "JoinModal",
  data: function() {
    return {
      roomId: ''
    }
  },
  methods: {
    testMethod: function() {
      this.$emit("myEvent", this.roomId);
    }
  }
}
</script>