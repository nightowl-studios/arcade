<template>
    <div class="word-choice-input">
        <b-row>
            <b-col class="word-choice-input__col text-center">
                <b-button
                    class="scribble-button"
                    variant="primary"
                    v-on:click="sendWord(0)"
                    >{{ words[0] }}</b-button
                >
            </b-col>
            <b-col class="word-choice-input__col text-center">
                <b-button
                    class="scribble-button"
                    variant="primary"
                    v-on:click="sendWord(1)"
                    >{{ words[1] }}</b-button
                >
            </b-col>
            <b-col class="word-choice-input__col text-center">
                <b-button
                    class="scribble-button"
                    variant="primary"
                    v-on:click="sendWord(2)"
                    >{{ words[2] }}</b-button
                >
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
        sendWord(index) {
            this.$gameApiService.selectWord(index);
        },
    },
};
</script>
