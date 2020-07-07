<template>
<div>
    <b-button v-b-modal.modal-2>Nickname</b-button>
    <b-modal id="modal-2" @ok="onOKClicked" title="Set your nickname">
        <input v-model="nickname" placeholder="Enter nickname">
    </b-modal>
</div>
</template>

<script>
export default {
    name: 'Nickname',
    data: function(){
        return{
            nickname: ''
        }
    },
    methods: {
        onOKClicked: function(){
            let setNickname = {
                nickname: this.nickname
            };
            this.$emit("onChangeNickname", setNickname)
        }
    }
}

</script>