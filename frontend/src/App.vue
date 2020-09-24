<template>
    <div id="app">
        <link
            href="https://fonts.googleapis.com/css2?family=Orbitron:wght@600&display=swap"
            rel="stylesheet"
        />
        <header class="app__header">
            <p v-on:click="goHome" class="app__header__logo">nightowl studios</p>
        </header>
        <div class="app__body">
            <router-view />
        </div>
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

<style lang="scss">
/* Move some of this stuff to index.scss */
* {
    padding:0;
    margin:0;
}

html, body {
    height:100%;
    display:flex;
    flex-direction:column;
}

body > * {
    flex-shrink: 0;
}

#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    height: 100%;
    display:flex;
    flex-direction:column;
}

.app__header {
    background-color: black;
    text-align: left;
    text-decoration: none;

    &__logo {
    display: inline-block;
    font-family: "Orbitron", sans-serif;
    padding: 5px;
    margin: 0;
    color: white;
    cursor: pointer;
    }
}

.app__body {
    flex-grow: 1;
}


</style>
