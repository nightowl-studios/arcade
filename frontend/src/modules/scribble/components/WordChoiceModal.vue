<template>
    <div class="word-choice-modal" v-if="words">
        <b-modal
            v-model="modalShow"
            ref="word-choice-modal"
            centered
            title="Choose a word to draw"
            hide-footer
            hide-header-close
            no-close-on-backdrop
            no-close-on-esc
        >
            <b-row>
                <b-col class="word-choice-modal__col text-center">
                    <WordChoiceInput :words="words" />
                </b-col>
            </b-row>
            <b-row>
                <b-col class="word-choice-modal__col text-center">
                    <BaseTimerBar />
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
    },

    components: {
        WordChoiceInput,
        BaseTimerBar,
    },
};
</script>
