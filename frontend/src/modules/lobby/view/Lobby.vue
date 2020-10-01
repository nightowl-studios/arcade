<template>
    <div id="lobby">
        <Header title="Sketch Night" />
        <Header class="lobby-header-room-id" :title="lobbyId" />
        <InvitationLink />
        <PlayerList class="player-list" :players="players" />
    </div>
</template>

<script>
import Header from "../components/Header.vue";
import PlayerList from "../components/PlayerList.vue";
import InvitationLink from "../components/InvitationLink.vue";
import { mapState } from "vuex";
import { EventBus } from "@/eventBus";
import { Event } from "@/events";

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
        playPlayerLeftSound() {
            let audio = new Audio(require("@/assets/audio/player-leave-2.mp3"));
            audio.play();
        },
        playPlayerJoinSound() {
            let audio = new Audio(require("@/assets/audio/player-join-2.mp3"));
            audio.play();
        },
    },
    created() {
        EventBus.$on(Event.PLAYER_LEFT, (data) => {
            console.log("PLAYER LEFTTTTTTT" + data);
            this.playPlayerLeftSound();
        });

        EventBus.$on(Event.PLAYER_JOIN, (data) => {
            console.log("PLAYER JOINNNN" + data);
            this.playPlayerJoinSound();
        });
    }
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
