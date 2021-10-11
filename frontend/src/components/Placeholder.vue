<template>
  <div class="block pl-5">
    <b-skeleton :animated="animated" class="py-3" :height="35" :width="400" />
    <b-skeleton :animated="animated" class="py-3" :height="200" :width="550" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  name: "Placeholder",
  props: {
    animated: Boolean,
  },
});
</script>
