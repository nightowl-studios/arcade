<template>
    <body>
        <div class="player">
            <h2 class="name">{{ name }}</h2>
        </div>
    </body>
</template>

<script>
export default {
    name: "Player",
    props: {
        name: String,
    },
};
</script>

<style scoped>
.player {
    border: solid 1px black;
    text-align: center;
    background-color: orange;
}
</style>
