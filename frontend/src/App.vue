<template>
    <div id="app">
        <link
            href="https://fonts.googleapis.com/css2?family=Orbitron:wght@600&display=swap"
            rel="stylesheet"
        />
        <header>
            <p v-on:click="goHome" class="header__logo">nightowl studios</p>
        </header>
        <router-view />
    </div>
</template>

<script>
export default {
    methods: {
        goHome: function () {
            this.$webSocketService.disconnect();
            this.$router.push({ name: "home" });
        },
    },
};
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    height: 100vh;
}

header {
    background-color: black;
    text-align: left;
    text-decoration: none;
}

.header__logo {
    display: inline-block;
    font-family: "Orbitron", sans-serif;
    padding: 5px;
    margin: 0;
    color: white;
    cursor: pointer;
}
</style>
