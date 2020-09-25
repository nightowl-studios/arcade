<template>
    <div class="header">
        <div class="header__message">
            <PlayerInfo
                v-if="gameState.showPlayerChoosing"
                :nickname="gameState.player.nickname"
            />
            <Word
                v-if="gameState.showWordToGuess"
                :word="gameState.word"
                :isGuessing="gameState.lockCanvas"
            />
        </div>
        <div class="header__timer">
            <BaseTimerCircle size="sm"/>
        </div>
    </div>
</template>

<script>
import Word from "../components/Word.vue";
import PlayerInfo from "../components/PlayerInfo.vue";
import BaseTimerCircle from "@/modules/common/components/BaseTimerCircle.vue";

import { mapState } from "vuex";
export default {
    components: {
        PlayerInfo,
        Word,
        BaseTimerCircle,
    },
    computed: {
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
    },
};
</script>

<style lang="scss" scoped>
.header {
    position: relative;
    height: 60px;

    &__message {
        position: absolute;
        top: 50%;
        left:50%;
        transform: translate(-50%, -50%);
        font-size: large;
        font-weight: bold;
    }

    &__timer {
        position: absolute;
        top: 5px;
        right: 5px;
    }
}
</style>

