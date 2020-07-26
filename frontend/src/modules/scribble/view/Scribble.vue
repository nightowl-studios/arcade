<template>
    <div class="scribble">
        <b-container fluid class="scribble__container">
            <WordChoice :words="gameState.words" />
            <b-row class="scribble__container__header" align-v="center">
                <b-col>
                    <Header :nickname="gameState.player.nickname" />
                </b-col>
            </b-row>
            <b-row class="scribble__container__body">
                <b-col>
                    <b-row class="scribble__container__body__players">
                        <PlayerList :players="players" />
                    </b-row>
                </b-col>
                <b-col>
                    <CanvasPanel
                        :colors="colors"
                        :sizes="sizes"
                        :isCanvasLocked="lockCanvas"
                    />
                </b-col>
                <b-col>
                    <b-row class="scribble__container__body__lobbyid">
                        <LobbyId />
                    </b-row>
                    <b-row class="scribble__container__body__chat">
                        <Chat />
                    </b-row>
                </b-col>
            </b-row>
        </b-container>
    </div>
</template>

<script>
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";
import Chat from "../components/Chat.vue";
import CanvasPanel from "../components/CanvasPanel.vue";
import Header from "../components/Header.vue";
import LobbyId from "../components/LobbyId.vue";
import PlayerList from "../components/PlayerList.vue";
import WordChoice from "../components/WordChoice.vue";
import { mapState } from "vuex";
import { WaitingForPlayerToChooseWord } from "../stores/states/gamestates";

export default {
    mixins: [WebSocketMixin],
    name: "Scribble",
    components: {
        CanvasPanel,
        Chat,
        Header,
        LobbyId,
        PlayerList,
        WordChoice,
    },
    data: function () {
        return {
            colors: ["#000000", "#4287f5", "#da42f5", "#7ef542"],
            sizes: [8, 16, 32, 64],
        };
    },
    computed: {
        ...mapState("application", {
            players: (state) => state.players,
        }),
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
        lockCanvas() {
            return this.gameState.state === WaitingForPlayerToChooseWord.STATE;
        },
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
            &__players {
                height: 50%;
                width: 30%;
                float: right;
            }

            &__lobbyid {
                margin-left: 25px;
            }

            &__chat {
                height: 50%;
                float: left;
            }
        }
    }
}
</style>
