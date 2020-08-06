<template>
    <div class="player-info">
        {{ nickname }} is currently selecting a word...
    </div>
</template>

<script>
export default {
    name: "PlayerInfo",
    props: {
        nickname: String,
    },
};
</script>

<style lang="scss" scoped>
.player-info {
    text-align: center;
}
</style>
