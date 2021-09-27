<template>
  <container>
    <Banner :title="title" :subtitle="subtitle" />
    <!-- <Cards :items="items" :perPage="12" color="light" /> -->
    <section>
      <b-field>
        <b-upload
          name="foo"
          v-model="fileUpload"
          accept="audio/*"
          drag-drop
          expanded
          type="is-primary"
        >
          <section class="section">
            <div class="content has-text-centered">
              <p>
                <b-icon icon="upload" size="is-large"> </b-icon>
              </p>
              <p>Drag and drop an audio clip</p>
            </div>
          </section>
        </b-upload>
      </b-field>
      <br />
      <div v-if="fileUpload.text" class="box block content">
        <p class="has-text-info">
          {{ fileUpload.name }}
        </p>
      </div>
    </section>
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
import { Player } from "tone";
import { Processor } from "@/lib/dsp";
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
    return {
      audioPlayer: Player,
      fileUpload: Blob,
    };
  },
  mounted() {
    this.updateFromDB();
  },
  methods: {
    ...mapMutations({
      updateFromDB: mu.Mutation(UPDATE_FROM_DB),
    }),
  },
  watch: {
    fileUpload: function(blob: Blob) {
      const player = new Player();
      const processor = new Processor(player, true)
        .WithFilter(400, "lowpass")
        .WithLimiter(-5);
      processor.Play(blob);
    },
  },
  computed: {
    ...mapState({
      items: (state: any) => state.clips.clips,
    }),
  },
});
</script>
