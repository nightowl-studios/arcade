<template>
    <span>
        <b-button class="lobby-button" variant="primary" v-b-modal.modal-2
            >Change Nickname</b-button
        >
        <b-modal
            ref="nicknameModal"
            id="modal-2"
            @ok="onOkClicked"
            title="Set your nickname"
        >
            <NicknameInput ref="nicknameInput" />
        </b-modal>
    </span>
</template>

<script>
import NicknameInput from "@/modules/common/components/NicknameInput.vue";

export default {
    name: "ChangeNicknameModal",

    components: {
        NicknameInput,
    },

    methods: {
        showModal() {
            this.$refs["nicknameModal"].show();
        },

        onOkClicked: function (nicknameModal) {
            nicknameModal.preventDefault();
            if (this.$refs["nicknameInput"].validateNickname()) {
                this.$refs["nicknameInput"].changeNickname();
                this.$nextTick(() => {
                    this.$refs["nicknameModal"].hide();
                });
            }
        },
    },
};
</script>
