<template>
  <Container>
      <Button @action="getMessage" text="Press me!"/>
      <Button text="Press me!"/>
      <Button text="Press me!"/>
      <Button text="Press me!"/>
      <Button text="Press me!"/>
  </Container>
</template>

<script>
import Container from "@/components/Container.vue"
import Button from "@/components/Button.vue"
export default {
  components: {
    Container,
    Button,
  },
  data() {
    return {
      message: "Press Me!"
    };
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic().then(result => {
        self.message = result;
        alert("hello");
      });
    }
  }
};
</script>