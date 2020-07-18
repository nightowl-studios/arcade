<template>
  <span>
    <b-button class="lobby-button" variant="primary" v-b-modal.modal-2>Change Nickname</b-button>
    <b-modal ref="nicknameModal" id="modal-2" @ok="onOKClicked" title="Set your nickname">
      <input v-model="nickname" placeholder="Enter nickname" />
    </b-modal>
  </span>
</template>

<script>
export default {
  name: "Nickname",
  data: function() {
    return {
      nickname: ""
    };
  },
  methods: {
    showModal() {
      this.$refs["nicknameModal"].show();
    },
    onOKClicked: function() {
      let message = {
        api: "hub",
        payload: {
          changeNameTo: this.nickname
        }
      };
      console.log("Sending msg");
      console.log(this.$webSocketService);
      this.$webSocketService.send(message);
    }
  },
  mounted() {
    //this.showModal();
  }
};
</script>
