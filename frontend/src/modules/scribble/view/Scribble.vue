<template>
    <div class="scribble">
        <Loading v-if="loading" />
        <Lobby v-else-if="gameState.showLobby" />
        <div v-else>
            <b-container fluid class="scribble__container">
                <b-row class="scribble__container__row">
                    <b-col class="scribble__container__col">
                        <LeftSidePanel/>
                    </b-col>
                    <b-col cols=6 class="scribble__container__col">
                        <CenterPanel />
                    </b-col>
                    <b-col class="scribble__container__col">
                        <RightSidePanel />
                    </b-col>
                </b-row>
                <Modal />
            </b-container>
            <b-modal
                :visible="gameState.showResults"
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
import LeftSidePanel from "../containers/LeftSidePanel.vue";
import Results from "../components/Results.vue";
import CenterPanel from "../containers/CenterPanel.vue";
import Modal from "../containers/Modal.vue";
import RightSidePanel from "../containers/RightSidePanel.vue";
import Lobby from "@/modules/lobby/view/Lobby.vue";
import Loading from "@/modules/common/view/Loading.vue";
import { mapGetters } from 'vuex'

export default {
    name: "Scribble",
    components: {
        LeftSidePanel,
        Lobby,
        CenterPanel,
        RightSidePanel,
        Modal,
        Results,
        Loading,
    },
    computed: {
        ...mapGetters("scribble", {
            loading: 'getLoading',
            gameState: 'getGameState'
        })
    },
    methods: {
        goHome() {
            // TODO: Reset vuex state
            this.$applicationController.closeWebSocket();
            this.$applicationController.resetStoreState();
            this.$scribbleGameController.resetStoreState();
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
    padding-top: 5%;
    padding-bottom: 5%;
    /* TODO only add margins when media size is large enough */
    padding-left: 10%;
    padding-right: 10%;

    &__container {
        height: 100%;

        &__row {
            height: 100%;
        }
    }

    &__col {
        margin: 2px;
        padding: 0;
    }
}
</style>
