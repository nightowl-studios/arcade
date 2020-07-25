<template>
    <div class="chatlog" ref="chatbox">
        <ul class="message">
            <li v-for="message in chatLog" :key="message.message">
                <p>
                    <span>{{ message[0] }}: {{ message[1] }}</span>
                </p>
            </li>
        </ul>
    </div>
</template>

<script>
import { EventBus } from "@/eventBus";
import { Event } from "@/events";

export default {
    name: "ChatLog",
    data: function() {
        return {
            newMessage: "",
            chatLog: [],
        };
    },
    created() {
        EventBus.$on(Event.CHAT_HISTORY, data => {
            for (const messages of data) {
                this.chatLog.push([messages.sender.nickname, messages.message]);
                this.$nextTick(() => {
                    let chatBox = this.$refs.chatbox;
                    if (chatBox) {
                        chatBox.scrollTop = chatBox.scrollHeight;
                    }
                });
            }
        });

        EventBus.$on(Event.CHAT_MESSAGE, data => {
            this.chatLog.push([data.sender.nickname, data.message]);
            this.$nextTick(() => {
                let chatBox = this.$refs.chatbox;
                if (chatBox) {
                    chatBox.scrollTop = chatBox.scrollHeight;
                }
            });
        });

        if (this.$webSocketService.isConnected()) {
            this.$chatApiService.requestChatHistory();
        } else {
            EventBus.$on(Event.WEBSOCKET_CONNECTED, () => {
                this.$chatApiService.requestChatHistory();
            });
        }
    },
};
</script>

<style scoped>
.chatlog {
    height: 88%;
    overflow: scroll;
    margin-bottom: 10px;
    border: 1px solid gray;
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
</style>
