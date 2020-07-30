<template>
    <div class="word-choice-modal" v-if="words">
        <b-modal
            v-model="modalShow"
            ref="word-choice-modal"
            centered
            title="Choose a word to draw"
            hide-footer
        >
            <WordChoiceInput :words="words" />
        </b-modal>
    </div>
</template>

<script>
import WordChoiceInput from "@/modules/scribble/components/WordChoiceInput.vue";
export default {
    name: "WordChoice",

    props: {
        words: Array,
        modalShow: Boolean,
    },
    components: {
        WordChoiceInput,
    },
};
</script>
