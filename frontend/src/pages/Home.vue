<template>
  <container>
    <Banner :title="title" :subtitle="subtitle" />
  </container>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "@/mixins/Page";
import Container from "@/components/Container.vue";
import Banner from "@/components/Banner.vue";

export default Vue.extend({
  name: "Home",
  mixins: [Page],
  components: {
    Container,
    Banner,
  },
});
</script>
