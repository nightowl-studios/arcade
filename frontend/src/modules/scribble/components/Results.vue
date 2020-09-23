<template>
    <div>
        <div class="results-list" v-for="player in players" :key="player.uuid">
            <p
                :style="{
                    visibility:
                        player.score === highestScore ? 'visible' : 'hidden',
                }"
            >
                👑
            </p>
            <p>{{ player.nickname }}</p>
            <p>{{ player.score }}pts</p>
        </div>
    </div>
</template>

<script>
export default {
    name: "Results",
    data: function () {
        return {
            players: [],
            highestScore: 0,
        };
    },
    created() {
        this.players = Object.create(this.$scribbleStoreService.getPlayers());
        this.players.sort(function (a, b) {
            if (a.score > b.score) return -1;
            if (a.score < b.score) return 1;
            return 0;
        });
        this.highestScore = this.players[0].score;
    },
};
</script>

<style scoped>
.results-list {
    display: grid;
    grid-template-columns: 2em 8em 2em;
}
</style>
