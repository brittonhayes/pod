<template>
  <Button
    @click="action"
    :type="fgColor"
    :size="size"
    :icon-left="icon"
    :class="[classes.button]"
    >{{ text }}</Button
  >
</template>

<script>
import Color from "@/mixins/Color.js";
import { Button } from "buefy/dist/components/button";
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
  components: { Button },
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
