<template>
    <div class="word-choice-input">
        <b-row>
            <b-col class="word-choice-input__col text-center">
                <b-button class="scribble-button" variant="primary">{{
                    words[0]
                }}</b-button>
            </b-col>
            <b-col class="word-choice-input__col text-center">
                <b-button class="scribble-button" variant="primary">{{
                    words[1]
                }}</b-button>
            </b-col>
            <b-col class="word-choice-input__col text-center">
                <b-button class="scribble-button" variant="primary">{{
                    words[2]
                }}</b-button>
            </b-col>
        </b-row>
        <!-- Test button to be removed when rick does this -->
        <b-row>
            <b-col class="text-center">
                <b-button v-on:click="sendWord">TEST</b-button>
            </b-col>
        </b-row>
    </div>
</template>

<script>
export default {
    name: "WordChoice",
    props: {
        words: Array,
    },
    methods: {
        sendWord() {
            // Select the first choice
            this.$gameApiService.selectWord(0);
        },
    },
};
</script>
