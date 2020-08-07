<template>
    <div class="base-timer" :class="baseTimerClasses">
        <svg
            class="base-timer__svg"
            viewBox="0 0 100 100"
            xmlns="http://www.w3.org/2000/svg"
        >
            <g class="base-timer__circle">
                <circle
                    class="base-timer__path-elapsed"
                    cx="50"
                    cy="50"
                    r="45"
                ></circle>
                <path
                    :stroke-dasharray="circleDasharray"
                    class="base-timer__path-remaining"
                    :style="transitionStyle"
                    :class="remainingPathColor"
                    d="M 50, 50 m -45, 0 a 45,45 0 1,0 90,0 a 45,45 0 1,0 -90,0"
                ></path>
            </g>
        </svg>
        <span class="base-timer__label" :class="baseTimerLabelClasses">
            {{ timeLeft }}
        </span>
    </div>
</template>

<script>
const FULL_DASH_ARRAY = 283;
const WARNING_THRESHOLD = 10;
const ALERT_THRESHOLD = 5;

const COLOR_CODES = {
    info: {
        color: "green",
    },
    warning: {
        color: "orange",
        threshold: WARNING_THRESHOLD,
    },
    alert: {
        color: "red",
        threshold: ALERT_THRESHOLD,
    },
};

export default {
    name: "BaseTimerCircle",

    props: {
        timeLimit: Number,
        size: String,
        state: String,
    },

    data() {
        return {
            timePassed: 0,
            timerInterval: null,
            transitionTime: "1s linear all",
        };
    },

    computed: {
        transitionStyle() {
            return {
                "--transitionTime-style": this.transitionTime,
            };
        },
        baseTimerClasses() {
            if (!this.size) {
                return `base-timer--md`;
            }
            return `base-timer--${this.size}`;
        },
        baseTimerLabelClasses() {
            if (!this.size) {
                return `base-timer__label--md`;
            }
            return `base-timer__label--${this.size}`;
        },
        circleDasharray() {
            return `${(this.timeFraction * FULL_DASH_ARRAY).toFixed(0)} 283`;
        },
        timeLeft() {
            return this.timeLimit - this.timePassed;
        },
        timeFraction() {
            const rawTimeFraction = this.timeLeft / this.timeLimit;
            return (
                rawTimeFraction - (1 / this.timeLimit) * (1 - rawTimeFraction)
            );
        },
        remainingPathColor() {
            const { alert, warning, info } = COLOR_CODES;

            if (this.timeLeft <= alert.threshold) {
                return alert.color;
            } else if (this.timeLeft <= warning.threshold) {
                return warning.color;
            } else {
                return info.color;
            }
        },
        watchState() {
            return this.state;
        },
        watchTimePassed() {
            return this.timePassed;
        },
    },

    watch: {
        timeLeft(newValue) {
            if (newValue === 0) {
                this.onTimesUp();
            }
        },
        watchState() {
            this.onTimesUp();
            this.transitionTime = "333ms linear all";
            this.startTimer();
        },
        watchTimePassed(newValue) {
            if (newValue > 0) {
                this.transitionTime = "1s linear all";
            }
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
    display: inline-block;
    position: relative;

    &--sm {
        width: 50px;
        height: 50px;
    }

    &--md {
        width: 100px;
        height: 100px;
    }

    &--lg {
        width: 200px;
        height: 200px;
    }

    &__svg {
        transform: scaleX(-1);
    }

    &__circle {
        fill: none;
        stroke: none;
    }

    &__path-elapsed {
        stroke-width: 7px;
        stroke: grey;
    }

    &__path-remaining {
        stroke-width: 7px;
        stroke-linecap: round;
        transform: rotate(90deg);
        transform-origin: center;
        transition: var(--transitionTime-style);
        fill-rule: nonzero;
        stroke: currentColor;

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

    &__label {
        position: absolute;
        top: 0;
        display: flex;
        align-items: center;
        justify-content: center;

        &--sm {
            font-size: 20px;
            width: 50px;
            height: 50px;
        }

        &--md {
            font-size: 48px;
            width: 100px;
            height: 100px;
        }

        &--lg {
            font-size: 48px;
            width: 200px;
            height: 200px;
        }
    }
}
</style>
