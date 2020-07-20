<template>
    <div id="scribble">
        <CanvasPanel :colors="colors" :sizes="sizes" />
        <Gameroom :clients="players" />
    </div>
</template>

<script>
import CanvasPanel from '../components/CanvasPanel.vue'
import Gameroom from '../components/Gameroom.vue'
import WebSocketMixin from '@/modules/common/mixins/webSocketMixin.js'

export default {
    mixins: [WebSocketMixin],
    name: 'Scribble',
    components: {
        CanvasPanel,
        Gameroom,
    },
    data: function () {
        return {
            colors: ['#000000', '#4287f5', '#da42f5', '#7ef542'],
            sizes: [8, 16, 32, 64],
        }
    },
}
</script>
