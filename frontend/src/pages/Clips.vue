<template>
  <container>
    <Banner :title="title" :subtitle="subtitle" />
    <!-- <Cards :items="items" :perPage="12" color="light" /> -->
    <div class="notification is-light has-text-centered">
      <h2 class="subtitle">No WAV files found</h2>
      <br />
      <b-tag type="is-primary" size="is-medium">~/pod/clips/*.wav</b-tag>
    </div>
  </container>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "@/mixins/Page";
import Container from "@/components/Container.vue";
import Banner from "@/components/Banner.vue";
// import Cards from "@/components/Cards.vue";

import { mapState, mapMutations } from "vuex";
import { Mutator, CLIPS, UPDATE_FROM_DB } from "@/store/mutations";

const mu = new Mutator(CLIPS);

export default Vue.extend({
  name: "Clips",
  mixins: [Page],
  components: {
    Banner,
    Container,
    // Cards,
  },
  data() {
    return {};
  },
  mounted() {
    this.updateFromDB();
  },
  methods: {
    ...mapMutations({
      updateFromDB: mu.Mutation(UPDATE_FROM_DB),
    }),
  },
  computed: {
    ...mapState({
      items: (state: any) => state.clips.clips,
    }),
  },
});
</script>
