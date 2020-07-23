<template>
    <div class="header">
        <h1>{{ title }}</h1>
    </div>
</template>

<script>
export default {
    name: "Header",
    props: {
        title: String,
    },
};
</script>

<style scoped>
.header {
    margin-top: 10px;
    margin-bottom: 10px;
}
</style>
