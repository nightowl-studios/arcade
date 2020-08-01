<template>
    <div>
        <h1 class="player-list__header">Players</h1>
        <div
            class="player-list__row"
            v-for="player in players"
            :key="player.uuid"
        >
            <p>[AVATAR]</p>
            <p>{{ player.nickname }}</p>
            <ChangeNicknameModal
                :style="{
                    visibility: isCurrentPlayer(player.uuid)
                        ? 'visible'
                        : 'hidden',
                }"
            />
        </div>
    </div>
</template>

<script>
import ChangeNicknameModal from "../components/ChangeNicknameModal.vue";
import { store } from "@/store";

export default {
    name: "PlayerList",
    props: {
        players: Array,
    },
    components: {
        ChangeNicknameModal,
    },
    methods: {
        isCurrentPlayer: function (playerUuid) {
            return store.getters["application/getPlayerUuid"] === playerUuid;
        },
    },
};
</script>

<style scoped>
.player-list__header {
    background-color: #111111;
    color: white;
    border: 1px solid black;
    margin: 0px;
    padding: 0.25em;
    border-radius: 5px 5px 0 0;
}

.player-list__row {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    justify-items: center;
    align-items: center;
    border: 1px solid black;
    border-top: 0px;
    padding: 0.5em;
}

.player-list__row:last-of-type {
    border-radius: 0 0 5px 5px;
}

.player-list__row p {
    margin: 0px;
}
</style>
