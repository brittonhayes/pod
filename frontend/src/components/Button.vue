<template>
    <b-button @click="action" :class="[padding,color]">{{text}}</b-button>
</template>

<script>
export default {
  data() {
    return {
      margin: "m-2",
      padding: "m-2"
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