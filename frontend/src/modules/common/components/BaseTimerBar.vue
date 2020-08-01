<template>
    <div class="base-timer">
        <svg
            class="base-timer__svg"
            viewBox="0 0 50 5"
            xmlns="http://www.w3.org/2000/svg"
        >
            <path class="base-timer__path-elapsed" d="M 1, 1 L 49, 1"></path>
            <path
                :stroke-dasharray="dasharray"
                class="base-timer__path-remaining"
                :class="remainingPathColor"
                d="M 1, 1 L 49, 1"
            ></path>
        </svg>
        <span class="base-timer__label">
            {{ timeLeft }}
        </span>
    </div>
</template>

<script>
const FULL_DASH_ARRAY = 48;
const WARNING_THRESHOLD = 5;
const DANGER_THRESHOLD = 2;
const COLOR_CODES = {
    healthy: {
        color: "green",
    },
    warning: {
        color: "orange",
        threshold: WARNING_THRESHOLD,
    },
    danger: {
        color: "red",
        threshold: DANGER_THRESHOLD,
    },
};
export default {
    name: "BaseTimerBar",
    props: {
        timeLimit: Number,
        size: String,
    },

    data() {
        return {
            timePassed: 0,
            timerInterval: null,
        };
    },
    computed: {
        dasharray() {
            return `${(this.timeFraction * FULL_DASH_ARRAY).toFixed(0)} 48`;
        },
        timeFraction() {
            return this.timeLeft / this.timeLimit;
        },
        timeLeft() {
            return this.timeLimit - this.timePassed;
        },
        watchTimeLimit() {
            return this.timeLimit;
        },
        remainingPathColor() {
            const { healthy, warning, danger } = COLOR_CODES;
            if (this.timeLeft <= DANGER_THRESHOLD) {
                return danger.color;
            } else if (this.timeLeft <= WARNING_THRESHOLD) {
                return warning.color;
            } else {
                return healthy.color;
            }
        },
    },
    watch: {
        timeLeft(newValue) {
            if (newValue === 0) {
                this.onTimesUp();
            }
        },
        watchTimeLimit() {
            this.onTimesUp();
            this.startTimer();
        },
    },
    mounted() {
        this.startTimer();
    },
    methods: {
        onTimesUp() {
            clearInterval(this.timerInterval);
            this.$emit("onTimesUp");
        },
        startTimer() {
            this.timePassed = 0;
            this.timerInterval = setInterval(
                () => (this.timePassed += 1),
                1000
            );
        },
    },
};
</script>

<style scoped lang="scss">
.base-timer {
    &__path-elapsed {
        stroke-width: 1px;
        stroke-linecap: round;
        stroke: grey;
    }
    &__path-remaining {
        stroke-width: 1px;
        stroke-linecap: round;
        transform-origin: right;
        transition: 1s linear all;
        stroke: currentColor;
        fill-rule: nonzero;
        &.green {
            color: rgb(65, 184, 131);
        }
        &.orange {
            color: orange;
        }
        &.red {
            color: red;
        }
    }
}
</style>
