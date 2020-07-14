<template>
<div>
    <div v-for="message in chatLog" :key="message">
        <h2 class="message">{{message[0]}}: {{message[1]}}</h2>
    </div>
</div>
</template>

<script>
import { EventBus } from '../eventBus';

export default {
    name: 'ChatLog',

    data: function(){
        return{
            chatLog: [],
        }
    },

    created() {
        EventBus.$on("chat", (data) => {
            if (data.history) {
                for (let messages of data.history){
                    this.chatLog.push([messages.sender.nickname, messages.message])
                }
            } else if (data.message) {
                this.chatLog.push([data.message.sender.nickname, data.message.message])
            }
        })

        let request = {
            "api":"chat",
            "payload":{
                "requestHistory": true
            }
        }
        this.$webSocketService.send(request)
    }
}
</script>

<style>
.message{
    font-size: 15px;
}
</style>