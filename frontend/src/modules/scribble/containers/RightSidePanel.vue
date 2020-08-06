<template>
    <div class="right-side-panel">
        <b-row class="right-side-panel__header">
            <RightSidePanelHeader />
        </b-row>
        <b-row class="right-side-panel__body">
            <Chat />
        </b-row>
    </div>
</template>

<script>
import RightSidePanelHeader from "./RightSidePanelHeader.vue";
import Chat from "../components/Chat.vue";

export default {
    components: {
        Chat,
        RightSidePanelHeader,
    },
};
</script>
