<template>
  <span>
    <b-button class="lobby-button" variant="primary" v-b-modal.modal-2>Change Nickname</b-button>
    <b-modal ref="nicknameModal" id="modal-2" @ok="onOKClicked" title="Set your nickname">
      <input  v-on:keyup.enter="onEnter" v-model="nickname" placeholder="Enter nickname" />
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
    hideModal(){
      this.$refs["nicknameModal"].hide();
    },
    onOKClicked: function() {
      let message = {
        api: "hub",
        payload: {
          changeNameTo: this.nickname
        }
      };
      this.$webSocketService.send(message);
    },
    onEnter: function(){
      let message = {
        api: "hub",
        payload: {
          changeNameTo: this.nickname
        }
      };
      this.$webSocketService.send(message);
      this.hideModal();
    }
  },
  mounted() {
    //this.showModal();
  }
};
</script>
