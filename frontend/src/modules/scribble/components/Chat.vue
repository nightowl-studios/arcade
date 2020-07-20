<template>
  <div id="wrapper">
    <div class="overflow" ref="chatbox">
      <ul class="message">
        <li v-for="message in chatLog" :key="message.message">
          <p>
            <span>{{message[0]}}: {{message[1]}}</span>
          </p>
        </li>
      </ul>
    </div>
    <div>
      <input
        v-on:keyup.enter="onSendMessage()"
        type="text"
        v-model="message"
        placeholder="Enter message"
      />
      <button v-on:click="onSendMessage()">Send</button>
    </div>
  </div>
</template>

<script>
import { createChatMessage } from "@/modules/common/utility/WebSocketMessageUtils";
import { EventBus } from "@/eventBus";

export default {
  name: "Chat",

  data: function() {
    return {
      newMessage: "",
      chatLog: [],
      message: ""
    };
  },
  created() {
    EventBus.$on("chat", data => {
      if (data.history) {
        for (let messages of data.history) {
          this.chatLog.push([messages.sender.nickname, messages.message]);
          this.$nextTick(() => {
            this.$refs["chatbox"].scrollTop = this.$refs[
              "chatbox"
            ].scrollHeight;
          });
        }
      } else if (data.message) {
        this.chatLog.push([data.message.sender.nickname, data.message.message]);
        this.$nextTick(() => {
          this.$refs["chatbox"].scrollTop = this.$refs["chatbox"].scrollHeight;
        });
      }
    });

    let request = {
      api: "chat",
      payload: {
        requestHistory: true
      }
    };
    this.$webSocketService.send(request);
  },

  methods: {
    onSendMessage: function() {
      if (this.message != "") {
        let messageToSend = createChatMessage(this.message);
        this.$webSocketService.send(messageToSend);
        this.message = "";
      }
    }
  }
};
</script>

<style>
#wrapper {
  margin: 0;
  padding-bottom: 10px;
  background: #e5f4fc;
  border: 1px solid #a1c5d8;
  display: inline-block;
}
.overflow {
  overflow: scroll;
  margin-bottom: 1px;
  margin: 10px;
  border: 1px solid gray;
  width: 400px;
  height: 500px;
  border-radius: 4px;
  overflow-x: hidden;
  display: flex;
}
.message {
  font-size: 15px;
  text-align: left;
  list-style-type: none;
  padding-left: 10px;
  padding-right: 10px;
  display: flex;
  flex-direction: column;
}
span {
  padding: 8px;
  border-radius: 4px;
}
p {
  float: left;
}
input {
  line-height: 3;
  width: 100%;
  border: 1px solid gray;
  border-left: none;
  border-bottom: none;
  border-right: none;
  border-bottom-left-radius: 4px;
  padding-left: 15px;
}
button {
  width: 145px;
  color: white;
  background: #4986b6;
  border-color: gray;
  border-bottom: none;
  border-right: none;
  border-bottom-right-radius: 3px;
}
</style>
