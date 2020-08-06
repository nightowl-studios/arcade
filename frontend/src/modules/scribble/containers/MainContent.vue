<template>
    <div>
        <CanvasPanel
            :colors="colors"
            :sizes="sizes"
            :isCanvasLocked="gameState.lockCanvas"
        />
    </div>
</template>

<script>
import CanvasPanel from "../components/CanvasPanel.vue";
import { mapState } from "vuex";

export default {
    components: {
        CanvasPanel,
    },
    data: function () {
        return {
            colors: ["#000000", "#4287f5", "#da42f5", "#7ef542", "#ffffff"],
            sizes: [8, 16, 32, 64],
        };
    },
    computed: {
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
    },
};
</script>
