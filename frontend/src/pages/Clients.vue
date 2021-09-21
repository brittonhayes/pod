<template>
  <div>
    <Banner :title="title" :subtitle="subtitle" />
    <Cards :items="items" color="light" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Banner from "@/components/Banner.vue";
import Cards from "@/components/Cards.vue";
import Page from "@/mixins/Page";
export default Vue.extend({
  mixins: [Page],
  components: {
    Banner,
    Cards,
  },
  data() {
    return {
      items: [
        { id: 1, name: "John Doe", summary: "" },
        { id: 2, name: "Jane Dosepth", summary: "" },
        { id: 3, name: "Sarah Marshal", summary: "" },
        { id: 4, name: "Hugh Bridges", summary: "" },
        { id: 5, name: "Alana Thomas", summary: "" },
        { id: 6, name: "Paula James", summary: "" },
        { id: 7, name: "Steve Rogers", summary: "" },
      ],
    };
  },
});
</script>
