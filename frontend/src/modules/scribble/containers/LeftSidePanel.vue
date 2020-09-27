<template>
    <b-container class="left-side-panel">
        <b-row>
            <b-col class="left-side-panel__col">
                <RoundNumber />
            </b-col>
        </b-row>
        <b-row>
            <b-col class="left-side-panel__col">
                <PlayerList :players="players" />
            </b-col>
        </b-row>
        <b-row>
            <b-col class="left-side-panel__col">
                <LobbyId />
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
import { mapState } from "vuex";
import PlayerList from "../components/PlayerList.vue";
import RoundNumber from "../components/RoundNumber.vue";
import LobbyId from "../components/LobbyId.vue";

export default {
    components: {
        PlayerList,
        RoundNumber,
        LobbyId,
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

    &__col {
        padding: 0;
        padding-bottom: 5px;
    }
}
</style>
