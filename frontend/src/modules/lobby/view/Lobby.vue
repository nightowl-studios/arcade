<template>
    <div id="lobby">
        <Header title="Sketch Night" />
        <Header class="lobby-header-room-id" :title="lobbyId" />
        <InvitationLink />
        <PlayerList class="player-list" :players="players" />
        <b-button v-on:click="requestGameInfo">Request</b-button>
    </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import InvitationLink from "../components/InvitationLink.vue";
import { mapState } from "vuex";

export default {
    name: "Lobby",

    components: {
        Header,
        PlayerList,
        InvitationLink,
    },

    computed: {
        ...mapState("application", {
            lobbyId: (state) => state.lobbyId,
        }),
        ...mapState("scribble", {
            players: (state) => state.players,
        }),
    },
    methods: {
        requestGameInfo: function () {
            console.log("requesting");
            this.$scribbleGameController.initGame();
        },
    },
};
</script>
<style scoped>
#lobby {
    display: grid;
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
