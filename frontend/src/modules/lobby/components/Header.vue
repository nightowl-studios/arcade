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
    margin: 0;
}

.header h1 {
    margin: 0;
}
</style>
