<template>
    <Button @click="action" :class="[margin,padding,color]">{{text}}</Button>
</template>

<script>
import {Button} from 'buefy/dist/components/button';
export default {
  components: {
      Button,
  },
  data() {
    return {
      margin: "m-2",
      padding: "p-2"
    };
  },
  methods: {
    action() {
        this.$emit("action")
    },
  },
  props: {
      color: {
        type: String,
        default: "is-primary",
        validator: function(value) {
            return ['is-primary', 'is-info', 'is-success', 'is-warning', 'is-danger', 'is-dark', 'is-light'].indexOf(value) !== -1
        },
      },
      text: {
          type: String, Number,
      },
  },
};
</script>