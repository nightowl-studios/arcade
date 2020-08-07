<template>
    <WordChoiceModal
        :words="gameState.words"
        :modalShow="gameState.showWordChoices"
        :timeLimit="gameState.durationSec"
    />
</template>

<script>
import { mapState } from "vuex";
import WordChoiceModal from "../components/WordChoiceModal.vue";

export default {
    components: {
        WordChoiceModal,
    },
    computed: {
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
    },
};
</script>
