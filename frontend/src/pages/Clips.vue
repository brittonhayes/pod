<template>
  <container>
    <Banner :title="title" :subtitle="subtitle" />
    <Cards :items="items" :perPage="12" color="light" />
  </container>
</template>

<script lang="ts">
import Page from "@/mixins/Page";
import Container from "@/components/Container.vue";
import Banner from "@/components/Banner.vue";
import Cards from "@/components/Cards.vue";
export default {
  name: "Clips",
  mixins: [Page],
  components: {
    Banner,
    Container,
    Cards,
  },
  data() {
    return {
      items: [
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
        {
          id: 1,
          name: "vo_mixdown_1.wav",
        },
      ],
    };
  },
};
</script>
