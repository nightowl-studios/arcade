<template>
    <div class="scribble">
        <b-container fluid class="scribble__container">
            <b-row class="scribble__container__row">
                <b-col class="scribble__container__col">
                    <LeftSide/>
                </b-col>
                <b-col cols=6 class="scribble__container__col">
                    <CenterPanel />
                </b-col>
                <b-col class="scribble__container__col">
                    <RightSide />
                </b-col>
            </b-row>
        </b-container>
    </div>
</template>

<script>
import LeftSide from "../containers/LeftSide.vue";
import CenterPanel from "../containers/CenterPanel.vue";
import RightSide from "../containers/RightSide.vue";
export default {
    components: {
        LeftSide,
        CenterPanel,
        RightSide,
    },
}
</script>

<style lang="scss" scoped>
.scribble {
    height: 100%;
    padding-top: 5%;
    padding-bottom: 5%;
    /* TODO only add margins when media size is large enough */
    padding-left: 10%;
    padding-right: 10%;

    &__container {
        height: 100%;

        &__row {
            height: 100%;
        }
    }

    &__col {
        margin: 2px;
        padding: 0;
    }
}
</style>
