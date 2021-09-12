<template>
  <div>
    <Banner :title="title" :subtitle="subtitle" />
    <Cards :items="items" color="light" />
  </div>
</template>

<script>
import Page from "@/mixins/Page.js";
import Banner from "@/components/Banner.vue";
import Cards from "@/components/Cards.vue";
export default {
  name: "Clips",
  mixins: [Page],
  components: {
    Banner,
    Cards,
  },
  data() {
    return {
      items: [
        {
          id: 1,
          title: "vo_mixdown_1.wav",
          body: "hello",
        },
      ],
    };
  },
};
</script>
