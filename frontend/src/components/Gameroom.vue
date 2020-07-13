<template>
<div>
    <h1 id="gameroom">Game room</h1>
    <div class="container">
        <div id="players">
            <h2>Players</h2>
            <div v-for="player in clients" :key="player">
                {{player.nickname}}
            </div>
        </div>
        <div id="drawing">
            <h3>Here is where the canvas goes</h3>
        </div>
        <div id="chatroom">
            <h2>Message Board</h2>
            <Chat/>
        </div>
    </div>
</div>
</template>


<script>

import Chat from './Chat.vue'

export default {
  name: 'Lobby',
  props: {
    clients: Array,
  },
  components: {
      Chat
  },
  data: function(){
      return {}
  },

}
</script>

<style>
#gameroom{
    text-align: center;
}
.container {
    display:flex;
    justify-content: space-between;
    width: 100%;
    height: 100%;
}

#players {
    flex: 1;
    padding-bottom: 100%;

    background-color:lightblue;
}
#drawing {
    flex: 2;
    padding-bottom: 100%;

    background-color: white;
}
#chatroom {
    flex:1;
    padding-bottom: 100%;
    background-color: lightblue;
}
</style>