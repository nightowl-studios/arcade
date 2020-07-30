<template>
    <div>
        <b-input-group>
            <b-form-input
                v-on:keyup.enter="onSendMessage()"
                type="text"
                v-model="message"
                placeholder="Enter message"
            ></b-form-input>
            <b-input-group-append>
                <b-button v-on:click="onSendMessage()" variant="info"
                    >Send</b-button
                >
            </b-input-group-append>
        </b-input-group>
    </div>
</template>

<script>
export default {
    name: "ChatInput",
    data: function () {
        return {
            message: "",
        };
    },
    methods: {
        onSendMessage: function () {
            console.log(`Sending message ${this.message}`);
            if (this.message !== "") {
                this.$chatApiService.sendChatMessage(this.message);
                this.message = "";
            }
        },
    },
};
</script>
