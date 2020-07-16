import { shallowMount } from '@vue/test-utils';
import Chat from '../../../src/modules/scribble/components/Chat.vue';

describe('Chat', () => {
    const wrapper = shallowMount(Chat)

    it('renders properly', () => {
        expect(wrapper.html()).toContain('<button>Send</button>');
    });
});
