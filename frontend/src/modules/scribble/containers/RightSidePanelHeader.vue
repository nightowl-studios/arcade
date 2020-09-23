<template>
    <div class="right-side-panel-header">
        <div class="right-side-panel-header__timer">
            <BaseTimerCircle size="sm"/>
        </div>
        <div class="right-side-panel-header__lobby-id">
            <LobbyId />
        </div>
    </div>
</template>

<script>
import LobbyId from "../components/LobbyId.vue";
import BaseTimerCircle from "@/modules/common/components/BaseTimerCircle.vue";

export default {
    components: {
        BaseTimerCircle,
        LobbyId,
    },
};
</script>
