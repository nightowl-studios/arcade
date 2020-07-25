<template>
    <div class="join-room">
        <b-form-input
            class="join-room__lobby-id-input"
            v-model="lobbyId"
            maxlength="4"
            placeholder="Lobby ID"
        ></b-form-input>
        <b-button variant="success" v-on:click="onJoin">
            Join
        </b-button>
        <p
            class="join-room__error"
            :style="{ visibility: showError ? 'visible' : 'hidden' }"
        >
            Lobby doesn't exist!
        </p>
    </div>
</template>

<script>
export default {
    name: "joinRoom",
    data: function () {
        return {
            lobbyId: "",
            showError: false,
        };
    },
    watch: {
        lobbyId: function (val) {
            this.lobbyId = val.toUpperCase();
            this.showError = false;
        },
    },
    methods: {
        onJoin: async function () {
            if (this.lobbyId.length !== 4) {
                this.showError = true;
                return;
            }
            const lobbyExists = await this.$hubApiService.checkLobbyExists(
                this.lobbyId
            );
            if (lobbyExists) {
                this.$emit("onJoinRoom", this.lobbyId);
            } else {
                this.showError = true;
            }
        },
    },
};
</script>

<style scoped>
.join-room {
    display: grid;
    grid-template-rows: auto;
    grid-gap: 7px;
}

.joinRoom * {
    width: 100%;
}

.join-room__lobby-id-input {
    text-align: center;
}

.join-room__error {
    display: inline-block;
    margin-top: 8px;
    color: red;
}
</style>
