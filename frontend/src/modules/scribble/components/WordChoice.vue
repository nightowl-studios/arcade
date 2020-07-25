<template>
    <div>
        <b-button
            class="scribble-button"
            variant="primary"
            v-b-modal.wordChoiceModalh
            >Choice</b-button
        >

        <b-modal
            ref="wordChoiceModal"
            id="wordChoiceModalh"
            title="Choose a word to draw"
            hide-footer
        >
            <WordChoiceInput ref="wordChoiceInput" :words="words" />
        </b-modal>
    </div>
</template>

<script>
import WordChoiceInput from "@/modules/scribble/components/WordChoiceInput.vue";
export default {
    name: "WordChoice",

    props: {
        words: Array,
    },
    components: {
        WordChoiceInput,
    },
};
</script>
