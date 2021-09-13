<template>
  <section :class="[classes.section, bgColor]">
    <div class="hero-body">
      <h1 :class="[classes.header, textColor]">
        {{ title }}
      </h1>
      <h2 class="subtitle is-size-6">
        {{ subtitle }}
      </h2>
    </div>
  </section>
</template>

<script>
import Text from "@/mixins/Text.js";
import Color from "@/mixins/Color.js";
export default {
  name: "Banner",
  mixins: [Color, Text],
  data() {
    return {
      classes: {
        header: {
          title: true,
          "is-size-4": true,
          "pb-2": true,
        },
        section: {
          "mb-3": true,
          hero: true,
          "is-small": true,
        },
      },
    };
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    subtitle: {
      type: String,
    },
  },
};
</script>
