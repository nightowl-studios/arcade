<template>
    <div class="right-side-panel-header">
        <div class="right-side-panel-header__timer">
            <BaseTimerCircle
                size="sm"
                :timeLimit="gameState.durationSec"
                :state="gameState.state"
            />
        </div>
        <div class="right-side-panel-header__lobby-id">
            <LobbyId />
        </div>
    </div>
</template>

<script>
import LobbyId from "../components/LobbyId.vue";
import BaseTimerCircle from "@/modules/common/components/BaseTimerCircle.vue";
import { mapState } from "vuex";

export default {
    components: {
        BaseTimerCircle,
        LobbyId,
    },
    computed: {
        ...mapState("scribble", {
            gameState: (state) => state.gameState,
        }),
    },
};
</script>
