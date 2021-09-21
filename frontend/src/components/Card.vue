<template>
  <div :class="[classes.card, CardColor]">
    <div class="card-content">
      <h3 :class="[classes.cardHeader, CardTextColor]">{{ title }}</h3>
      <div :class="[classes.cardContent, CardTextColor]">
        <p>
          {{ body }}
        </p>
        <slot />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { MustBeOneOf } from "@/lib/validate";
import { Color, BackgroundColor, TextColor, Palette } from "@/mixins/Color";
export default Vue.extend({
  name: "Card",
  data() {
    return {
      classes: {
        card: {
          card: true,
          column: true,
          "p-2": true,
        },
        cardHeader: {
          "has-text-weight-bold": true,
          "is-size-6": true,
          "pb-5": true,
        },
        cardContent: {
          content: true,
        },
      },
    };
  },
  props: {
    title: {
      type: String,
      required: false,
    },
    body: {
      type: String,
      required: false,
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
    CardColor: function() {
      return BackgroundColor(this.color as Color);
    },
    CardTextColor: function() {
      return TextColor(this.color as Color);
    },
  },
});
</script>
