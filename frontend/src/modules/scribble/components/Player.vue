<template>
    <b-row class="player">
        <b-col>{{ player.nickname }}</b-col>
        <b-col>{{ player.score }}</b-col>
    </b-row>
</template>

<script>
export default {
    name: "Player",
    props: {
        player: Object,
    },
};
</script>
