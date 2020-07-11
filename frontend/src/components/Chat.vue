<template>
<div>
    <input type="text" v-model="message" placeholder="Enter message">
    <button v-on:click="onSendMessage">Send</button>
    <ChatLog/>
</div>
</template>

<script>
import { ArcadeWebSocket } from '../webSocket'
import ChatLog from './ChatLog'

export default {
    name: 'Chat',

    components: {
        ChatLog
    },
    data: function(){
        return{
            newMessage: '',
            chatLog: [],
            message: ''
        }
    },
    methods: {
      onSendMessage: function() {
          let messageToSend = {
              "api":"chat",
              "payload":{
                  "message": this.message
              }
          }
          ArcadeWebSocket.send(messageToSend);
    },
}
}
</script>