<template>
    <div>
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
