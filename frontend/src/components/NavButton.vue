<template>
  <b-button
    :to="to"
    tag="router-link"
    :size="size"
    :class="classObject"
    :type="ButtonColor"
    :icon-left="icon"
  />
</template>

<script lang="ts">
import Vue from "vue";
import { MustBeOneOf } from "@/lib/validate";
import { Color, Palette, IsColor } from "@/mixins/Color";

export default Vue.extend({
  name: "NavButton",
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
    color: {
      type: String,
      default: Color.White,
      validator: (value: Color) => {
        return MustBeOneOf("color", value, Palette);
      },
    },
  },
  computed: {
    ButtonColor: function() {
      return IsColor(this.color as Color);
    },
    classObject: function() {
      return {
        "m-2": true,
      };
    },
  },
});
</script>
