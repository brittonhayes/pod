<template>
  <div>
    <Banner :title="title" :subtitle="subtitle" />
    <section class="section">
      <div class="columns is-multiline">
        <card color="success" class="p-0">
          <h1 class="is-size-2 p-1 has-text-centered">🌴</h1>
        </card>
        <card color="success" class="p-0">
          <h1 class="is-size-2 p-1 has-text-centered">🌴</h1>
        </card>
        <card color="success" class="p-0">
          <h1 class="is-size-2 p-1 has-text-centered">🌴</h1>
        </card>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "@/mixins/Page";
import Banner from "@/components/Banner.vue";
import Card from "@/components/Card.vue";

export default Vue.extend({
  name: "Home",
  mixins: [Page],
  components: {
    Banner,
    Card,
  },
  data() {
    return {
      media: 3,
    };
  },
});
</script>
