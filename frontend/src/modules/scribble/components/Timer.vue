<template>
    <div class="timer">{{ timeLeft }}</div>
</template>

<script>
export default {
    name: "Timer",
    data() {
        return {
            timerInterval: null,
            timePassed: 0,
        };
    },
    props: {
        startTime: Number,
    },
    computed: {
        timeLeft() {
            return this.timeInSeconds(this.startTime) - this.timePassed;
        },
    },
    methods: {
        startTimer() {
            this.timerInterval = setInterval(
                () => (this.timePassed += 1),
                1000
            );
        },
        timeInSeconds(nanoseconds) {
            const nanoSecondToSecondsFactor = 1000000000;
            return nanoseconds / nanoSecondToSecondsFactor;
        },
    },
    mounted() {
        this.startTimer();
    },
};
</script>
