<template>
    <div class="round-number">
        <b-container fluid>
            <b-row align-v="center" class="justify-content-md-center">
                <b-col md="auto">Round: </b-col>
                <b-col md="auto"
                    ><input
                        class="round-number__input"
                        type="text"
                        v-model="roundNum"
                        readonly
                /></b-col>
            </b-row>
        </b-container>
    </div>
</template>

<script>
export default {
    name: "RoundNumber",

    props: {
        roundNumber: Number,
    },

    data: function () {
        return {
            roundNum: this.roundNumber + 1,
        };
    },
};
</script>

<style lang="scss" scoped>
.round-number {
    background-color: #e5f4fc;
    border: 1px solid #a1c5d8;
    padding-top: 5px;
    padding-bottom: 5px;
    border-radius: 6px;
    margin: 5px 5px 5px 5px;
}

.round-number__input {
    color: black;
    width: 60px;
    text-align: center;
    border-radius: 6px;
}
</style>
