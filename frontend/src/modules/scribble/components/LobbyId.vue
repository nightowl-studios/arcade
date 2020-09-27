<template>
    <div class="lobby-id">
        Lobby Id: {{ lobbyId }}
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
}

</style>
