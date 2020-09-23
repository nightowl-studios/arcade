<template>
    <div class="right-side-panel-header">
        <div class="right-side-panel-header__timer">
            <BaseTimerCircle size="sm"/>
        </div>
    </div>
</template>

<script>
import BaseTimerCircle from "@/modules/common/components/BaseTimerCircle.vue";

export default {
    components: {
        BaseTimerCircle,
    },
};
</script>
