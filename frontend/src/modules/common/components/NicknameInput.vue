<template>
    <div id="nicknameInput">
        <b-form-input v-model="nickname" placeholder="Enter nickname" />
        <p
            id="invalidNicknameError"
            :style="{ visibility: this.showError ? 'visible' : 'hidden' }"
        >
            Please enter a valid nickname!
        </p>
    </div>
</template>

<script>
import { createChangeNicknameMessage } from "@/utility/WebSocketMessageUtils";

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
p {
    margin: 5px;
}

#invalidNicknameError {
    margin-top: 8px;
    color: red;
}
</style>
