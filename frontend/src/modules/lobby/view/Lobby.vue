<template>
    <div id="lobby">
        <Header title="Sketch Night" />
        <Header class="lobby-header-room-id" :title="lobbyId" />
        <InvitationLink />
        <div class="lobby-buttons">
            <b-button
                class="lobby-button"
                variant="success"
                v-on:click="startGame"
                >Start game</b-button
            >
            <b-button
                class="exit-button"
                variant="danger"
                v-on:click="exitToHome"
                >Exit Lobby</b-button
            >
        </div>
        <PlayerList class="player-list" :players="players" />
    </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import InvitationLink from "../components/InvitationLink.vue";
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";
import { Event } from "@/events";
import { EventBus } from "@/eventBus.js";

export default {
    name: "Lobby",

    components: {
        Header,
        PlayerList,
        InvitationLink,
    },

    data: function () {
        return {
            lobbyId: "",
        };
    },

    mixins: [WebSocketMixin],

    methods: {
        startGame: function () {
            this.$gameApiService.startGame();
        },

        exitToHome: function () {
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
    display: grid;
    justify-items: center;
    grid-gap: 1em;
}

.lobby-buttons {
    display: grid;
    grid-template-columns: auto auto;
    justify-items: center;
    grid-gap: 1em;
}

.lobby-header-room-id {
    color: orange;
}

.player-list {
    width: 700px;
}
</style>
