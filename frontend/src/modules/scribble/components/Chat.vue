<template>
    <div id="chat">
        <ChatLog />
        <ChatInput />
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

<style scoped>
#chat {
    margin: 0;
    padding: 10px;
    background: #e5f4fc;
    border: 1px solid #a1c5d8;
    width: 100%;
}
</style>
