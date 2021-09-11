import { mount } from '@vue/test-utils'
import Contact from '@/components/Contact.vue'

describe('Contact.vue', () => {
  it('renders a contact info card', () => {
    const contact = {
        name: "John",
        email: "john@doe.com",
        description: "A cool guy"
    }
    const wrapper = mount(Contact, {
      propsData: { 
          name: contact.name,
          email: contact.email,
          description: contact.description,
        }
    })
    
    expect(wrapper.text()).toContain(contact.name)
    expect(wrapper.text()).toContain(contact.email)
    expect(wrapper.text()).toContain(contact.description)
  })
})
