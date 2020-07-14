<template>
  <div>
    <input type="text" v-model="message" placeholder="Enter message" />
    <button v-on:click="onSendMessage">Send</button>
    <ChatLog />
  </div>
</template>

<script>
import { createChatMessage } from "../utility/WebSocketMessageUtils";
import ChatLog from "./ChatLog";

export default {
  name: "Chat",

  components: {
    ChatLog
  },
  data: function() {
    return {
      newMessage: "",
      chatLog: [],
      message: ""
    };
  },
  methods: {
    onSendMessage: function() {
      let messageToSend = createChatMessage(this.message);
      this.$webSocketService.send(messageToSend);
    }
  }
};
</script>
