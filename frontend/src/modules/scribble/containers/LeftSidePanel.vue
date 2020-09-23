<template>
    <div class="left-side-panel">
        <PlayerList :players="players" />
    </div>
</template>

<script>
import { mapState } from "vuex";
import PlayerList from "../components/PlayerList.vue";

export default {
    components: {
        PlayerList,
    },
    computed: {
        ...mapState("scribble", {
            players: (state) => state.players,
        }),
    },
};
</script>

<style lang="scss" scoped>
.left-side-panel {
    height: 100%;
    border: 1px black solid;
}
</style>
