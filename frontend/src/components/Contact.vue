<template>
  <div class="">
    <p>{{ name }}</p>
    <p>{{ email }}</p>
    <p>{{ description }}</p>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  props: {
    name: {
      type: String,
      required: true,
    },
    email: {
      type: String,
    },
    description: {
      type: String,
    },
  },
});
</script>
