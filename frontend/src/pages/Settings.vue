<template>
  <div>
    <Banner :title="title" :subtitle="subtitle" />
  </div>
</template>

<script lang="ts">
import Page from "@/mixins/Page";
import Banner from "@/components/Banner.vue";
export default {
  mixins: [Page],
  components: {
    Banner,
  },
};
</script>
