<template>
    <b-container class="left-side-panel">
        <b-row>
            <b-col>
                <RoundNumber />
            </b-col>
        </b-row>
        <b-row>
            <b-col>
            <PlayerList :players="players" />
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
import { mapState } from "vuex";
import PlayerList from "../components/PlayerList.vue";
import RoundNumber from "../components/RoundNumber.vue";

export default {
    components: {
        PlayerList,
        RoundNumber,
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
