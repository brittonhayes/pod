<template>
  <b-button
    :to="to"
    tag="router-link"
    :size="size"
    :class="classObject"
    :type="fgColor"
    :icon-left="icon"
  />
</template>

<script>
import Color from "@/mixins/Color.js";
export default {
  data() {
    return {};
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
  computed: {
    classObject: function() {
      return {
        "m-2": true,
      };
    },
  },
};
</script>
