<template>
  <b-button
    :to="to"
    tag="router-link"
    :size="size"
    class="m-2"
    :type="fgColor"
    :icon-left="icon"
  />
</template>

<script>
import Color from "@/mixins/Color.js";
export default {
  data() {
    return {
      margin: "m-2",
    };
  },
  mixins: [Color],
  props: {
    to: {
      type: String,
    },
    icon: {
      type: String,
      default: "user",
    },
    size: {
      type: String,
      default: "is-medium",
    },
  },
};
</script>
