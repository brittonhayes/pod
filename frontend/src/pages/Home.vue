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
    <section class="mx-5">
      <article v-for="i in media" :key="i">
        <feed class="notification is-light">
          <template v-slot:header>
            <h1>{{ i }}</h1>
          </template>
          <template v-slot:subtitle>
            <p>Here might be a page title</p>
          </template>
          <template v-slot:icon-left>
            <p>x</p>
          </template>
          <template v-slot:icon-right>
            <p>o</p>
          </template>
        </feed>
      </article>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "../mixins/Page";
import Banner from "../components/Banner.vue";
import Card from "../components/Card.vue";
import Feed from "../components/Feed.vue";

export default Vue.extend({
  name: "Home",
  mixins: [Page],
  components: {
    Banner,
    Card,
    Feed,
  },
  data() {
    return {
      media: 3,
    };
  },
});
</script>
