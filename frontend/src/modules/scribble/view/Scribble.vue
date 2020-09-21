<template>
    <div class="scribble">
        <Loading v-if="loading" />
        <Lobby v-else-if="gameState.showLobby" />
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
            <b-modal
                v-model="showResults"
                @ok="goHome"
                title="Results"
                centered
                ok-only
                size="lg"
                no-close-on-esc
                no-close-on-backdrop
                hide-header-close
            >
                <Results></Results>
            </b-modal>
        </div>
    </div>
</template>

<script>
import Header from "../containers/Header.vue";
import LeftSidePanel from "../containers/LeftSidePanel.vue";
import Results from "../components/Results.vue";
import MainContent from "../containers/MainContent.vue";
import Modal from "../containers/Modal.vue";
import RightSidePanel from "../containers/RightSidePanel.vue";
import Lobby from "@/modules/lobby/view/Lobby.vue";
import { mapState } from "vuex";
import Loading from "@/modules/common/view/Loading.vue";

export default {
    name: "Scribble",
    components: {
        Header,
        LeftSidePanel,
        Lobby,
        MainContent,
        RightSidePanel,
        Modal,
        Results,
        Loading,
    },
    computed: {
        ...mapState("scribble", {
            loading: (state) => state.loading,
            gameState: (state) => state.gameState,
            showResults: (state) => state.gameState.showResults,
        }),
    },
    methods: {
        goHome() {
            this.$webSocketService.disconnect();
            this.$router.push({ name: "home" });
        },
    },
    async created() {
        const lobbyId = this.$router.currentRoute.params.lobbyId;
        const goodToGo = await this.$applicationController.initApplication(lobbyId);
        if (!goodToGo) {
            console.log("Room is invalid");
        }
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
