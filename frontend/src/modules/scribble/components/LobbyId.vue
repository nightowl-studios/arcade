<template>
    <div class="lobby-id">
        <b-container fluid>
            <b-row align-v="center" class="justify-content-md-center">
                <b-col md="auto">Lobby ID: </b-col>
                <b-col md="auto"
                    ><input
                        class="lobby-id__input"
                        type="text"
                        v-model="lobbyId"
                        readonly
                /></b-col>
            </b-row>
        </b-container>
    </div>
</template>

<script>
export default {
    name: "LobbyId",
    data: function () {
        return {
            lobbyId: window.location.href.split("/").pop(),
        };
    },
};
</script>

<style lang="scss" scoped>
.lobby-id {
    background-color: #e5f4fc;
    border: 1px solid #a1c5d8;
    padding-top: 5px;
    padding-bottom: 5px;
    border-radius: 6px;
    margin: 5px 5px 5px 5px;
}

.lobby-id__input {
    color: black;
    width: 60px;
    text-align: center;
    border-radius: 6px;
}
</style>
