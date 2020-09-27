<template>
    <b-container fluid class="center-panel">
        <b-row class="center-panel__header" align-v="center">
            <b-col class="center-panel__col">
                <Header />
            </b-col>
        </b-row>
        <b-row class="center-panel__canvas">
            <CanvasPanel
                :colors="colors"
                :sizes="sizes"
                :isCanvasLocked="gameState.lockCanvas"
            />
        </b-row>
    </b-container>
</template>

<script>
import Header from "../containers/Header.vue";
import CanvasPanel from "../components/CanvasPanel.vue";
import { mapState } from "vuex";

export default {
    components: {
        Header,
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


<style lang="scss" scoped>
.center-panel {
    height: 100%;
    border: 1px black solid;

    &__col {
        padding: 0;
    }
}
</style>
