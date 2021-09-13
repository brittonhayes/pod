<template>
  <div>
    <Banner :title="title" :subtitle="subtitle" />
  </div>
</template>

<script>
import Page from "@/mixins/Page.js";
import Banner from "@/components/Banner.vue";
export default {
  mixins: [Page],
  components: {
    Banner,
  },
};
</script>
