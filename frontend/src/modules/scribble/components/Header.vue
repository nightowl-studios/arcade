<template>
    <div class="header">{{ nickname }} is currently drawing...</div>
</template>

<script>
export default {
    name: "Header",
    props: {
        nickname: String,
    },
}
</script>

<style lang="scss" scoped>
.header {
    text-align: center;
}
</style>

