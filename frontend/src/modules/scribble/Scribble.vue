<template>
    <div class="scribble">
        <Lobby v-if="gameState.showLobby" />
        <div v-else>
            <b-container fluid class="scribble__container">
                <b-row class="scribble__container__header" align-v="center">
                    <b-col>
                        <Header />
                    </b-col>
                </b-row>
                <b-row class="scribble__container__body">
                    <b-col>
                        <LeftSidePanel />
                    </b-col>
                    <b-col>
                        <MainContent />
                    </b-col>
                    <b-col>
                        <RightSidePanel />
                    </b-col>
                </b-row>
                <Modal />
            </b-container>
        </div>
    </div>
</template>

<script>
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";
import Header from "./containers/Header.vue";
import LeftSidePanel from "./containers/LeftSidePanel.vue";
import MainContent from "./containers/MainContent.vue";
import Modal from "./containers/Modal.vue";
import RightSidePanel from "./containers/RightSidePanel.vue";
import Lobby from "@/modules/lobby/view/Lobby.vue";
import { mapState } from "vuex";
import { WaitingInLobby } from "@/modules/scribble/stores/states/gamestates";

export default {
    mixins: [WebSocketMixin],
    name: "Scribble",
    components: {
        Header,
        LeftSidePanel,
        Lobby,
        MainContent,
        RightSidePanel,
        Modal,
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

<style lang="scss" scoped>
.scribble {
    height: 100%;

    &__container {
        height: 100%;

        &__body {
            height: 100%;
        }
    }
}
</style>
