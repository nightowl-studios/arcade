<template>
  <div class="invitation-link-section">
    <b-container fluid>
      <b-row align-v="center" class="justify-content-md-center">
        <b-col md="auto">Invitation Link: </b-col>
        <b-col md="auto"><input id="invitation-link" type="text" v-model="inviteLink" readonly/></b-col>
        <b-col md="auto"><b-button class="copy-link-button" variant="success" v-on:click="copyInviteLink">Copy Link</b-button></b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
export default {
  name: "InvitationLink",
  data: function() {
    return {
      inviteLink: window.location.href
    };
  },
  methods: {
    copyInviteLink: function() {
      let linkToCopy = document.querySelector("#invitation-link");
      linkToCopy.select();
      try {
        var successful = document.execCommand("copy");
        var msg = successful ? "successful" : "unsuccessful";
        alert("Invitation link was copied " + msg);
      } catch (err) {
        alert("Oops, unable to copy");
      }
    }
  }
};
</script>

<style scoped>
.invitation-link-section {
  margin-top: 25px;
}

#invitation-link {
  color: black;
  width: 300px;
}
</style>
