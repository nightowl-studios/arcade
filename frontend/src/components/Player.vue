<template>
<body>
    <div class = "player">
        <h2 class="name">{{name}}</h2>
        <h2 class="id">{{id}}</h2>
    </div>
</body>
</template>

<script>
export default{
    name: 'Player',
    props: {
        name: String,
        id: String
    },

    methods: {
        checkExistingID: function (id,name){
            var existingID = [];
            if (!existingID.includes(id)){
                console.log( name + "have joined the lobby.")
                existingID.push(id);
            }
        }
    }
}
</script>

<style scoped>
    body {
        background-image: url('images/scribbleLobby.jpg');
        background-repeat: no-repeat;
        background-size: cover;
        background-position: center;
    }
    h2{
        font-size: 12px;
    }
    .name {
        color:black;
        padding-right: 20%;
    }
    .id {
        color: black;
    }
    .player {
        display:flex;
        justify-content: center;
        padding: 5px;
        background-color:tomato;
        margin-left: 20%;
        margin-right: 20%;
        margin-bottom: 5px;
        border :2px solid black;
    }
</style>