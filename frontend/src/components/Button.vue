<template>
  <b-button
    @click="action"
    :type="fgColor"
    :size="size"
    :icon-left="icon"
    :class="[classes.button]"
    >{{ text }}</b-button
  >
</template>

<script>
import Color from "@/mixins/Color.js";
export default {
  data() {
    return {
      classes: {
        button: {
          margin: "m-2",
          padding: "p-2",
        },
      },
    };
  },
  mixins: [Color],
  methods: {
    action() {
      this.$emit("action");
    },
  },
  props: {
    text: {
      type: [String, Number],
    },
    icon: {
      type: String,
      default: "",
    },
    size: {
      type: String,
      default: "",
    },
  },
};
</script>
