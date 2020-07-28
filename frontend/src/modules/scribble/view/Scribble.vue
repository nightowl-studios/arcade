<template>
    <div class="scribble">
        <b-container fluid class="scribble__container">
            <b-row class="scribble__container__header" align-v="center">
                <b-col>
                    <Header
                        v-if="gameState.showPlayerChoosing"
                        nickname="gameState.player.nickname"
                    />
                    <Word
                        v-if="gameState.showWordToGuess"
                        :word="gameState.word"
                        :isGuessing="gameState.lockCanvas"
                    />
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
                        :isCanvasLocked="gameState.lockCanvas"
                    />
                </b-col>
                <b-col>
                    <b-row class="scribble__container__body__lobbyid">
                        <LobbyId />
                    </b-row>
                    <b-row class="scribble__container__body__chat">
                        <Chat :sendToScribbleApi="isGuessing" />
                    </b-row>
                </b-col>
            </b-row>
            <WordChoiceModal
                :words="gameState.words"
                :modalShow="gameState.showWordChoices"
            />
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
import WordChoiceModal from "../components/WordChoiceModal.vue";
import { mapState } from "vuex";
import { Guessing } from "../stores/states/gamestates";
import Word from "../components/Word.vue";

export default {
    mixins: [WebSocketMixin],
    name: "Scribble",
    components: {
        CanvasPanel,
        Chat,
        Header,
        LobbyId,
        PlayerList,
        WordChoiceModal,
        Word,
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
        isGuessing() {
            return this.gameState.state === Guessing.STATE;
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
