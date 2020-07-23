<template>
    <div id="scribble">
        <b-container fluid class="scribble-container">
            <b-row class="scribble-row-header">{{ chosenPlayer }} is drawing...</b-row>
            <b-row class="scribble-row-main">
                <b-col>
                    <CanvasPanel :colors="colors" :sizes="sizes" />
                </b-col>
                <b-col>
                    <b-row class="scribble-row-players">
                        <PlayerList :players="players" />
                    </b-row>
                    <b-row class="scribble-row-chat">
                        <Chat />
                    </b-row>
                </b-col>
            </b-row>
        </b-container>
    </div>
</template>

<script>
import CanvasPanel from "../components/CanvasPanel.vue";
import WebSocketMixin from "@/modules/common/mixins/webSocketMixin.js";
import Chat from "../components/Chat.vue";
import PlayerList from "../components/PlayerList.vue";
import { mapState } from "vuex";

export default {
    mixins: [WebSocketMixin],
    name: "Scribble",
    components: {
        CanvasPanel,
        Chat,
        PlayerList,
    },
    data: function() {
        return {
            colors: ["#000000", "#4287f5", "#da42f5", "#7ef542"],
            sizes: [8, 16, 32, 64],
        };
    },
    computed: {
        ...mapState('scribble', {
            chosenUuid: state => state.chosenUuid,
        }),
        ...mapState('application', {
            players: state => state.players
        }),
        chosenPlayer() {
            const chosenPlayer = this.players.filter(player => player.uuid === player.uuid)[0];
            return chosenPlayer.nickname;
        }
    },
};
</script>

<style scoped>
#scribble {
    height: 100%;
}

.scribble-container {
    height: 100%;
}

.scribble-row-main {
    height: 100%;
}

.scribble-row-players {
    height: 50%;
}

.scribble-row-chat {
    height: 50%;
}
</style>
