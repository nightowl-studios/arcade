<template>
    <div class="timer">{{ timeLeft }}</div>
</template>

<script>
export default {
    name: "Timer",

    props: {
        timeLeft: Number,
    },
};
</script>
