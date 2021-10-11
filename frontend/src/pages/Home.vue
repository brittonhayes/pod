<template>
  <container>
    <Banner :title="title" :subtitle="subtitle" :construction="true" />
    <placeholder animated />
  </container>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "@/mixins/Page";
import Container from "@/components/Container.vue";
import Banner from "@/components/Banner.vue";
import Placeholder from "@/components/Placeholder.vue";

export default Vue.extend({
  name: "Home",
  mixins: [Page],
  components: {
    Container,
    Banner,
    Placeholder,
  },
});
</script>
