<template>
  <container>
    <banner :title="title" :subtitle="subtitle">
      <template slot="buttons">
        <div class="block buttons is-hidden-mobile">
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
      <b-modal
        :active="modalEnabled"
        has-modal-card
        :width="1500"
        :can-cancel="true"
        @close="toggleModal"
      >
        <project-form v-bind="project"></project-form>
      </b-modal>
    </section>
    <div v-if="projects.length > 0">
      <project-card
        v-for="project in projects"
        :key="project.id"
        :project="project.name"
        :client="project.client"
        :content="project.summary"
      />
    </div>
    <div v-else class="notification has-text-centered py-6">
      <p class="subtitle">{{ empty }}</p>
    </div>
  </container>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import Page from "@/mixins/Page";
import Container from "@/components/Container.vue";
import Banner from "@/components/Banner.vue";
import ProjectForm from "@/components/ProjectForm.vue";
import ProjectCard from "@/components/ProjectCard.vue";

import { mapGetters, mapMutations, mapState } from "vuex";
import {
  Mutator,
  IS_ENABLED,
  SET_ENABLED,
  TOGGLE_ENABLED,
  UPDATE_FROM_DB,
  PROJECTS,
} from "@/store/mutations";
import { Project } from "@/types/project";

const mu = new Mutator(PROJECTS);

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
      empty: "No projects",
      project: {
        type: Object as PropType<Project>,
        default: () => {},
      },
    };
  },
  methods: {
    ...mapMutations({
      updateFromDB: mu.Mutation(UPDATE_FROM_DB),
      toggleModal: mu.Mutation(TOGGLE_ENABLED),
    }),
  },
  beforeMount() {
    this.$store.commit(mu.Mutation(SET_ENABLED), false);
  },
  mounted() {
    this.$store.commit(mu.Mutation(UPDATE_FROM_DB));
  },
  computed: {
    ...mapGetters({ modalEnabled: mu.Mutation(IS_ENABLED) }),
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
