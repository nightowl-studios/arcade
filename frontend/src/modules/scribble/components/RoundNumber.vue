<template>
    <div class="round-number">
        <span> Round {{ round }} </span>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
    name: "RoundNumber",
    computed: {
        ...mapGetters("scribble", {
            round: 'getRoundNumber',
        })
    },
};
</script>

<style lang="scss" scoped>
.round-number {
    border: 1px solid #a1c5d8;
}
</style>
