<template>
    <div class="scribble">
        <Lobby v-if="gameState.showLobby" />
    </div>
</template>

<script>
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";

import Lobby from "@/modules/lobby/view/Lobby.vue";
import { mapState } from "vuex";
import { WaitingInLobby } from "@/modules/scribble/stores/states/gamestates";

export default {
    mixins: [WebSocketMixin],
    name: "Scribble",
    components: {
        Lobby,
    },
    computed: {
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
    },
    created() {
        this.$store.commit("application/setLobbyId", this.lobbyId);

        const startState = new WaitingInLobby();
        this.$store.commit("scribble/setGameState", startState);
    },
};
</script>

<style scoped></style>
