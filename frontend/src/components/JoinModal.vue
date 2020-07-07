<template>
  <div id="joinModal">
    <b-button variant="success" v-b-modal.modal-1>Join</b-button>
    <b-modal id="modal-1" @ok="onOKClicked" title="BootstrapVue">
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

<style scoped>
#joinModal {
  margin-top: 15px;
}
</style>
