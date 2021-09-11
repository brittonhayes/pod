import { mount } from '@vue/test-utils'
import Button from '@/components/Button.vue'

describe('Button.vue', () => {
  it('renders a button with custom text', () => {
    const text = 'Click me!'
    const wrapper = mount(Button, {
      propsData: { text }
    })
    expect(wrapper.text()).toMatch(text)
  })
})
