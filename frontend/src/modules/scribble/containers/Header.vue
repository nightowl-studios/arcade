<template>
    <div>
        <PlayerInfo
            v-if="gameState.showPlayerChoosing"
            :nickname="gameState.player.nickname"
        />
        <Word
            v-if="gameState.showWordToGuess"
            :word="gameState.word"
            :isGuessing="gameState.lockCanvas"
        />
        <BaseTimerCircle style="float: right" size="sm"/>
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
