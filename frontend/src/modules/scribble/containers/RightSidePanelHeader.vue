<template>
    <div class="right-side-panel-header">
        <div class="right-side-panel-header__timer">
            <BaseTimerCircle size="sm" :timeLimit="gameState.duration" />
        </div>
        <div class="right-side-panel-header__lobby-id">
            <LobbyId />
        </div>
        <div class="right-side-panel-header__roundNumber">
            <RoundNumber />
        </div>
    </div>
</template>

<script>
import LobbyId from "../components/LobbyId.vue";
import BaseTimerCircle from "@/modules/common/components/BaseTimerCircle.vue";
import RoundNumber from "../components/RoundNumber.vue";
import { mapState } from "vuex";

export default {
    components: {
        BaseTimerCircle,
        LobbyId,
        RoundNumber,
    },
    computed: {
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
    },
};
</script>
