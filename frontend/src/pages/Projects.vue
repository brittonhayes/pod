<template>
  <container>
    <banner :title="title" :subtitle="subtitle">
      <template slot="buttons">
        <div class="block buttons">
          <b-button
            icon-left="plus"
            label="Create project"
            type="is-success is-light"
            @click="toggleModal"
          />
          <b-button icon-left="redo" type="is-light" @click="updateFromDB" />
        </div>
      </template>
    </banner>
    <section>
      <b-modal :active="modalEnabled" has-modal-card :width="1500">
        <project-form v-bind="project"></project-form>
      </b-modal>
    </section>
    <project-card
      v-for="project in projects"
      :key="project.id"
      :project="project.name"
      :client="project.client"
      :content="project.summary"
    />
  </container>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "@/mixins/Page";
import Container from "@/components/Container.vue";
import Banner from "@/components/Banner.vue";
import ProjectForm from "@/components/ProjectForm.vue";
import ProjectCard from "@/components/ProjectCard.vue";

import { mapGetters, mapMutations, mapState } from "vuex";
import {
  IS_ENABLED,
  SET_ENABLED,
  TOGGLE_ENABLED,
  UPDATE_FROM_DB,
} from "@/store/modules/mutations";

export default Vue.extend({
  mixins: [Page],
  components: {
    Banner,
    Container,
    ProjectForm,
    ProjectCard,
  },
  data() {
    return {
      project: {
        name: "",
        summary: "",
        client: "",
      },
    };
  },
  methods: {
    ...mapMutations({
      updateFromDB: UPDATE_FROM_DB,
      toggleModal: TOGGLE_ENABLED,
    }),
  },
  beforeMount() {
    this.$store.commit(SET_ENABLED, false);
  },
  mounted() {
    this.updateFromDB();
  },
  computed: {
    ...mapGetters({ modalEnabled: IS_ENABLED }),
    ...mapState({
      projects: (state: any) => state.projects.projects,
    }),
    classObject: function() {
      return {
        panel: true,
      };
    },
  },
});
</script>
