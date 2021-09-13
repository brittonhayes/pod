<template>
  <div
    :class="[
      classes.tile,
      isAncestor,
      isChild,
      isNotification,
      isVertical,
      hasMargins,
    ]"
  >
    <h2 v-if="title != ''" class="is-size-6">{{ title }}</h2>
    <slot />
  </div>
</template>

<script>
import { IsStyle } from "@/lib/validate.js";
export default {
  name: "Tile",
  data() {
    return {
      classes: {
        tile: {
          tile: true,
        },
      },
    };
  },
  computed: {
    hasMargins: function() {
      if (this.margin > 0) {
        return `m-${this.margin}`;
      }
      return 0;
    },
    isNotification: function() {
      if (this.notification) {
        return "notification";
      }
      return "";
    },
    isAncestor: function() {
      return IsStyle(this.ancestor, "ancestor");
    },
    isParent: function() {
      return IsStyle(this.parent, "parent");
    },
    isVertical: function() {
      return IsStyle(this.vertical, "vertical");
    },
    isChild: function() {
      return IsStyle(this.child, "child");
    },
  },
  props: {
    title: {
      type: String,
      default: "",
    },
    margin: {
      type: Number,
      default: 1,
    },
    notification: {
      type: Boolean,
    },
    ancestor: {
      type: Boolean,
    },
    parent: {
      type: Boolean,
    },
    child: {
      type: Boolean,
    },
    vertical: {
      type: Boolean,
    },
  },
};
</script>
