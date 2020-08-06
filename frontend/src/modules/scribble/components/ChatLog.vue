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
    data: function () {
        return {
            newMessage: "",
            chatLog: [],
        };
    },
    methods: {
        playCorrectSound() {
            let audio = new Audio(require("@/assets/audio/sound-correct.wav"));
            audio.play();
        },
        playPlayerLeftSound() {
            let audio = new Audio(require("@/assets/audio/player-leave.wav"));
            audio.play();
        },
    },
    created() {
        EventBus.$on(Event.CHAT_HISTORY, (data) => {
            for (const messages of data) {
                this.chatLog.push([messages.sender.nickname, messages.message]);
                this.$nextTick(() => adjustScrollTop());
            }
        });

        EventBus.$on(Event.CHAT_MESSAGE, (data) => {
            this.chatLog.push([data.sender.nickname, data.message]);
            this.$nextTick(() => adjustScrollTop());
        });

        EventBus.$on(Event.CORRECT_GUESS, (data) => {
            this.chatLog.push([data.nickname, "guessed correctly"]);
            this.playCorrectSound();
            this.$nextTick(() => adjustScrollTop());
        });

        EventBus.$on(Event.PLAYER_LEFT, (data) => {
            this.chatLog.push([data.nickname, "has left the game"]);
            this.playPlayerLeftSound();
            this.$nextTick(() => adjustScrollTop());
        });

        EventBus.$on(Event.PLAYER_JOIN, (data) => {
            this.chatLog.push([data.nickname, "has joined the game"]);
            this.playPlayerLeftSound();
            this.$nextTick(() => adjustScrollTop());
        });

        if (this.$webSocketService.isConnected()) {
            this.$chatApiService.requestChatHistory();
        } else {
            EventBus.$on(Event.WEBSOCKET_CONNECTED, () => {
                this.$chatApiService.requestChatHistory();
            });
        }

        function adjustScrollTop() {
            const chatBox = this.$refs.chatbox;
            if (chatBox) {
                chatBox.scrollTop = chatBox.scrollHeight;
            }
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
