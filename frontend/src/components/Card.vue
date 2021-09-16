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

<script>
import Color from "@/mixins/Color.js";
export default {
  name: "Card",
  mixins: [Color],
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
  },
};
</script>
