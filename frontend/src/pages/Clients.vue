<template>
  <div>
    <Banner :title="title" :subtitle="subtitle" color="light" />
    <Cards :items="items" color="success" />
  </div>
</template>

<script>
import Page from "@/mixins/Page.js";
import Banner from "@/components/Banner.vue";
import Cards from "@/components/Cards.vue";
export default {
  mixins: [Page],
  components: {
    Banner,
    Cards,
  },
  data() {
    return {
      items: [
        { id: 1, title: "John Doe", body: "" },
        { id: 2, title: "Jane Dosepth", body: "" },
        { id: 3, title: "Sarah Marshal", body: "" },
        { id: 4, title: "Hugh Bridges", body: "" },
        { id: 5, title: "Alana Thomas", body: "" },
        { id: 6, title: "Paula James", body: "" },
        { id: 7, title: "Steve Rogers", body: "" },
      ],
    };
  },
};
</script>
