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

<script lang="ts">
import Vue from "vue";
import { MustBeOneOf } from "@/lib/validation";
import { Color, Palette, IsColor } from "@/mixins/Color";
export default Vue.extend({
  name: "Button",
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
  },
});
</script>
