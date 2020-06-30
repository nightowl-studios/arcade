<template>
  <div>
    <b-button v-b-modal.modal-1>Launch demo modal</b-button>

    <b-modal id="modal-1" @ok="onOKClicked" title="BootstrapVue">
      <p class="my-4">Hello from modal!</p>
      <input v-model="hubId" placeholder="Enter room id">
    </b-modal>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "JoinModal",
  data: function() {
    return {
      hubId: ''
    }
  },
  methods: {
    onOKClicked: function() {
      console.log("Joining room " + this.hubId + "...")
      let apiUrl = this.$httpURL + '/hub' + '/' + this.hubId;
      axios
        .get(apiUrl)
        .then(response => {
          let responseWithHubId = {
            hubId: this.hubId,
            response: response
          };
          this.$emit("onJoinRoom", responseWithHubId)
        });
    }
  }
}
</script>