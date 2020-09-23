<template>
    <div class="right-side-panel">
        <Chat />
    </div>
</template>

<script>
import Chat from "../components/Chat.vue";

export default {
    components: {
        Chat,
    },
};
</script>

<style lang="scss" scoped>
.right-side-panel {
    height: 100%;
    border: 1px black solid;
}
</style>
