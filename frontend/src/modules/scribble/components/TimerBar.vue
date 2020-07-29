<template>
    <div class="base-timer">
        <svg
            class="base-timer__svg"
            viewBox="0 0 50 5"
            xmlns="http://www.w3.org/2000/svg"
        >
            <path
                class="base-timer__path-elapsed"
                d="
            M 1, 1
            L 49, 1
          "
            ></path>
            <path
                :stroke-dasharray="Dasharray"
                class="base-timer__path-remaining"
                :class="remainingPathColor"
                d="
            M 1, 1
            L 49, 1
          "
            ></path>
        </svg>
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
    name: "BarTimer",

    props: {
        timeLeft: Number,
        timeLimit: Number,
    },

    computed: {
        Dasharray() {
            return `${(this.timeFraction * FULL_DASH_ARRAY).toFixed(0)} 48`;
        },
        timeFraction() {
            return this.timeLeft / this.timeLimit;
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
