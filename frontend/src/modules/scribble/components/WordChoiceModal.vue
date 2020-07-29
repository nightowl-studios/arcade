<template>
    <div class="word-choice-modal">
        <b-modal
            v-model="modalShow"
            ref="word-choice-modal"
            centered
            title="Choose a word to draw"
            hide-footer
        >
            <b-row class="mb-4">
                <b-col class="word-choice-modal__col text-center">
                    <WordChoiceInput :words="words" />
                </b-col>
            </b-row>
            <b-row>
                <b-col class="word-choice-modal__col text-center">
                    <TimerBar
                        :timeLeft="timer.timeLeft()"
                        :timeLimit="timer.timeLimit"
                    />
                </b-col>
            </b-row>
            <b-row>
                <b-col class="word-choice-modal__col text-center">
                    <TimerCountdown :timeLeft="timer.timeLeft()" />
                </b-col>
            </b-row>
        </b-modal>
    </div>
</template>

<script>
const TIME_LIMIT = 10;

import WordChoiceInput from "@/modules/scribble/components/WordChoiceInput.vue";
import TimerBar from "@/modules/scribble/components/TimerBar.vue";
import TimerCountdown from "@/modules/scribble/components/TimerCountdown.vue";
import Timer from "@/utility/Timer.js";
export default {
    name: "WordChoice",

    props: {
        words: Array,
        modalShow: Boolean,
    },

    components: {
        WordChoiceInput,
        TimerBar,
        TimerCountdown,
    },

    data() {
        return {
            timer: null,
        };
    },

    created() {
        this.timer = new Timer(TIME_LIMIT);
        this.timer.startTimer();
    },

    computed: {
        timeLeft() {
            return this.timer.timeLeft();
        },
    },

    watch: {
        timeLeft(newValue) {
            console.log("Hi");
            if (newValue === 0) {
                this.timer.onTimesUp();
            }
        },
    },
};
</script>
