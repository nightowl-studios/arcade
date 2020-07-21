<template>
    <div id="lobby">
        <Header title="Welcome to Not Scribble" />
        <Header class="lobby-header-room-id" :title="lobbyId" />
        <div class="lobby-buttons">
            <b-button class="lobby-button" variant="success" v-on:click="startGame">Start game</b-button>
            <ChangeNicknameModal />
            <b-button class="exit-button" variant="danger" v-on:click="exitToHome">Exit Lobby</b-button>
        </div>
        <PlayerList :players="players" />
    </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import ChangeNicknameModal from "../components/ChangeNicknameModal.vue";
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";
import { Event } from "@/events";
import { EventBus } from "@/eventBus.js";

export default {
    name: "Lobby",

    components: {
        Header,
        PlayerList,
        ChangeNicknameModal,
    },

    data: function() {
        return {
            lobbyId: "",
        };
    },

    mixins: [WebSocketMixin],

    methods: {
        startGame: function() {
            this.$gameApiService.startGame();
        },

        exitToHome: function() {
            this.$webSocketService.disconnect();
            this.$router.push({ name: "home" });
        },
    },
    created() {
        EventBus.$on(Event.START_GAME, () => {
            this.$router.push({ path: `/scribble/${this.lobbyId}` });
        });
    },
};
</script>
<style scoped>
#lobby {
    padding-left: 100px;
    padding-right: 100px;
}

.lobby-button,
.exit-button {
    margin-left: 2px;
    margin-right: 2px;
}

.lobby-header-room-id {
    color: orange;
}
</style>
