<template>
  <div :class="[classes.card, bgColor]">
    <div class="card-content">
      <h3 :class="[classes.cardHeader, textColor]">{{ title }}</h3>
      <div :class="[classes.cardContent, textColor]">
        {{ body }}
        <slot />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { MustBeOneOf } from "../lib/validate";
import { Color, BackgroundColor, TextColor, Palette } from "../mixins/Color";
export default Vue.extend({
  name: "Card",
  data() {
    return {
      classes: {
        card: {
          card: true,
          column: true,
          "m-2": true,
        },
        cardHeader: {
          "has-text-weight-bold": true,
          "is-size-6": true,
          "pb-5": true,
        },
        cardContent: {
          content: true,
          "is-size-7": true,
        },
      },
    };
  },
  props: {
    title: {
      type: String,
    },
    body: {
      type: String,
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
    BannerColor: function() {
      return BackgroundColor(this.color as Color);
    },
    BannerTextColor: function() {
      return TextColor(this.color as Color);
    },
  },
});
</script>
