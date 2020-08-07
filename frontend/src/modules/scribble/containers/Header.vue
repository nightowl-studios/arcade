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
    </div>
</template>

<script>
import Word from "../components/Word.vue";
import PlayerInfo from "../components/PlayerInfo.vue";

import { mapState } from "vuex";
export default {
    components: {
        PlayerInfo,
        Word,
    },
    computed: {
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
    },
};
</script>
