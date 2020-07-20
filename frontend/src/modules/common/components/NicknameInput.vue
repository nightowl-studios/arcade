<template>
    <div id="nicknameInput">
        <input v-model="nickname" placeholder="Enter nickname" />
        <p
            id="invalidNicknameError"
            :style="{ visibility: this.showError ? 'visible' : 'hidden' }"
        >
            Please enter a valid nickname!
        </p>
    </div>
</template>

<script>
import { createChangeNicknameMessage } from "@/modules/common/utility/WebSocketMessageUtils";

export default {
    name: "NicknameInput",

    data: function () {
        return {
            nickname: "",
            showError: false,
        };
    },

    watch: {
        nickname: function () {
            this.showError = false;
        },
    },

    methods: {
        validateNickname: function () {
            const validNickname = "".localeCompare(this.nickname) !== 0;
            this.showError = !validNickname;
            return validNickname;
        },

        changeNickname: function () {
            this.$webSocketService.send(
                createChangeNicknameMessage(this.nickname)
            );
        },
    },
};
</script>
<style scoped>
input {
    border: 1px;
    border-style: solid;
    border-radius: 5px;
    border-color: grey;
}

p {
    margin: 5px;
}

#invalidNicknameError {
    margin-top: 8px;
    color: red;
}
</style>
