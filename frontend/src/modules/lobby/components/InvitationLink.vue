<template>
    <div class="invitation-link">
        <p class="invitation-link__title">Invitation Link</p>
        <input
            class="invitation-link__url"
            type="text"
            v-model="inviteLink"
            readonly
        />
        <b-button
            class="invitation-link__button"
            variant="primary"
            v-on:click="copyInviteLink"
            >Copy</b-button
        >
    </div>
</template>

<script>
export default {
    name: "InvitationLink",
    data: function () {
        return {
            inviteLink: window.location.href,
        };
    },
    methods: {
        copyInviteLink: function () {
            const linkToCopy = document.querySelector(".invitation-link__url");
            linkToCopy.select();
            try {
                var successful = document.execCommand("copy");
                if (successful) {
                    alert(`Invitation link copied!`);
                } else {
                    alert("Oops, unable to copy");
                }
            } catch (err) {
                alert("Oops, unable to copy");
            }
        },
    },
};
</script>

<style scoped>
.invitation-link {
    display: grid;
    grid-template-columns: auto 300px auto;
    grid-template-rows: 38px;
    border-radius: 10px;
    border: 1px solid black;
}

.invitation-link__title {
    margin: 0;
    padding-left: 10px;
    padding-right: 10px;

    display: flex;
    align-items: center;
}

.invitation-link__url {
    color: black;
    padding: 10px;
    border: 0;
    border-left: 1px solid black;
    border-right: 1px solid black;
}

.invitation-link__button {
    border-radius: 0 9px 9px 0;
}
</style>
