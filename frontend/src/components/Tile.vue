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

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
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
      if (this.ancestor) {
        return "is-ancestor";
      }
      return "";
    },
    isParent: function() {
      if (this.parent) {
        return "is-parent";
      }
      return "";
    },
    isVertical: function() {
      if (this.vertical) {
        return "is-vertical";
      }
      return "";
    },
    isChild: function() {
      if (this.child) {
        return "is-child";
      }
      return "";
    },
  },
});
</script>
