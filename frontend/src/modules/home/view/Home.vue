<template>
    <div id="home">
        <Title id="title" msg="SketchNight" />
        <NicknameInput ref="nicknameInput" />
        <div class="start-game">
            <CreateButton @onCreateRoom="onCreateRoom" />
            <JoinRoom @onJoinRoom="onJoinRoom" />
        </div>
    </div>
</template>

<script>
import Title from "../components/Title.vue";
import CreateButton from "../components/CreateButton.vue";
import JoinRoom from "../components/JoinRoom.vue";
import NicknameInput from "@/modules/common/components/NicknameInput.vue";

export default {
    name: "Home",

    components: {
        Title,
        CreateButton,
        JoinRoom,
        NicknameInput,
    },

    methods: {
        onCreateRoom: function (lobbyId) {
            this.connectToRoom(lobbyId);
        },

        onJoinRoom: function (lobbyId) {
            this.connectToRoom(lobbyId);
        },

        connectToRoom: function (lobbyId) {
            if (!this.$refs["nicknameInput"].validateNickname()) {
                return;
            }
            this.$router.push({ name: "lobby", params: { lobbyId: lobbyId } });
            this.$applicationController.setLobbyId(lobbyId);
            this.$refs["nicknameInput"].changeNickname();
        },
    },
};
</script>

<style scoped>
#home {
    margin-top: 2em;
    display: grid;
    grid-gap: 1em;
    grid-template-rows: auto;
    justify-items: center;
}

.start-game {
    display: grid;
    grid-gap: 5em;
    grid-template-columns: 200px 200px;
    justify-content: center;
    align-content: center;
}

/*
    Resources to follow
    - https://codepen.io/byron-seto/pen/MWyqNKO
    - https://codepen.io/byron-seto/pen/eYZLqBr

*/
</style>
