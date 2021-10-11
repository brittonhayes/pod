<template>
  <container>
    <Banner :title="title" :subtitle="subtitle" construction />
    <placeholder animated />
  </container>
</template>

<script lang="ts">
import Page from "@/mixins/Page";
import Banner from "@/components/Banner.vue";
import Container from "@/components/Container.vue";
import Placeholder from "@/components/Placeholder.vue";
export default {
  mixins: [Page],
  components: {
    Banner,
    Container,
    Placeholder,
  },
};
</script>
