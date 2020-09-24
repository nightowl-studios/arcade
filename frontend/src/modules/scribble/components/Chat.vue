<template>
    <div class="chat">
        <div class="chat__log">
            <ChatLog />
        </div>
        <div class="chat__input">
            <ChatInput />
        </div>
    </div>
</template>

<script>
import ChatInput from "./ChatInput.vue";
import ChatLog from "./ChatLog.vue";

export default {
    name: "Chat",
    components: {
        ChatInput,
        ChatLog,
    },
};
</script>

<style lang="scss" scoped>
.chat {
    height: 100%;
    position: relative;
    background: #e5f4fc;

    &__log {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 95%;
    }

    &__input {
        width: 100%;
        position: absolute;
        bottom: 5px;
    }
}
</style>
