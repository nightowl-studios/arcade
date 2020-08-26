<template>
    <div>
        <b-button
            class="create-button"
            variant="success"
            v-on:click="createRoom"
            >Create Room</b-button
        >
    </div>
</template>

<script>
export default {
    name: "CreateButton",
    methods: {
        createRoom: async function () {
            let lobbyId = await this.$apiSenderFacade.createLobby();
            this.$emit("onCreateRoom", lobbyId);
        },
    },
};
</script>

<style scoped>
.create-button {
    width: 100%;
}
</style>
