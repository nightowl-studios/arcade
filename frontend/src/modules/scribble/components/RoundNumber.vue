<template>
    <div class="round-number">
        <span> Round {{ roundNumber }} </span>
    </div>
</template>

<script>
export default {
    name: "RoundNumber",
    props: {
        roundNumber: Number,
    },
};
</script>

<style lang="scss" scoped>
.round-number {
    background-color: #e5f4fc;
    border: 1px solid #a1c5d8;
    padding-top: 5px;
    padding-bottom: 5px;
    border-radius: 6px;
    margin: 5px auto 5px auto;
    width: 90px;
}
</style>
