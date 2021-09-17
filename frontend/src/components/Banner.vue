<template>
  <section :class="[classes.section, BannerColor]">
    <div class="hero-body">
      <h1 :class="[classes.header, BannerTextColor]">
        {{ title }}
      </h1>
      <h2 :class="[classes.subheader, BannerTextColor]">
        {{ subtitle }}
      </h2>
    </div>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { MustBeOneOf } from "../lib/validate";
import { Color, Palette, BackgroundColor, TextColor } from "../mixins/Color";
export default Vue.extend({
  name: "Banner",
  data() {
    return {
      classes: {
        header: {
          title: true,
          "is-size-4": true,
          "pb-2": true,
        },
        subheader: {
          subtitle: true,
          "is-size-6": true,
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
