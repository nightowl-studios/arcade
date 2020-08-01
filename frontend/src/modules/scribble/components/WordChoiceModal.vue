<template>
    <div class="word-choice-modal" v-if="words">
        <b-modal
            v-model="modalShow"
            ref="word-choice-modal"
            centered
            title="Choose a word to draw"
            hide-footer
        >
            <b-row>
                <b-col class="word-choice-modal__col text-center">
                    <WordChoiceInput :words="words" />
                </b-col>
            </b-row>
            <b-row>
                <b-col class="word-choice-modal__col text-center">
                    <BaseTimerBar :timeLimit="timeLimit" />
                </b-col>
            </b-row>
        </b-modal>
    </div>
</template>

<script>
import WordChoiceInput from "@/modules/scribble/components/WordChoiceInput.vue";
import BaseTimerBar from "@/modules/common/components/BaseTimerBar.vue";
export default {
    name: "WordChoice",

    props: {
        words: Array,
        modalShow: Boolean,
        timeLimit: Number,
    },

    components: {
        WordChoiceInput,
        BaseTimerBar,
    },
};
</script>
