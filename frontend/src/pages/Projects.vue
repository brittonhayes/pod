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
    <Table
      v-if="items"
      :data="items"
      :columns="columns"
      :sort="sort"
      :selected="project"
    />
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
import Table from "@/components/Table.vue";

import { mapGetters, mapMutations, mapState } from "vuex";
import {
  IS_ENABLED,
  Namespace,
  SET_ENABLED,
  TOGGLE_ENABLED,
  UPDATE_FROM_DB,
} from "@/store/mutations";
import { Project, ProjectColumns, ProjectSort } from "@/types/project";

export default Vue.extend({
  mixins: [Page],
  components: {
    Banner,
    Container,
    Table,
    ProjectForm,
  },
  data() {
    return {
      empty: "No projects",
      project: {
        type: Object as PropType<Project>,
        default: () => {},
      },
      sort: ProjectSort,
      columns: ProjectColumns,
    };
  },
  methods: {
    ...mapMutations({
      updateFromDB: Namespace.Projects + UPDATE_FROM_DB,
      toggleModal: Namespace.Projects + TOGGLE_ENABLED,
    }),
  },
  beforeMount() {
    this.$store.commit(Namespace.Projects + SET_ENABLED, false);
  },
  mounted() {
    this.$store.commit(Namespace.Projects + UPDATE_FROM_DB);
  },
  computed: {
    ...mapGetters({ modalEnabled: Namespace.Projects + IS_ENABLED }),
    ...mapState({
      items: (state: any) => state.projects.projects,
    }),
    classObject: function() {
      return {
        panel: true,
      };
    },
  },
});
</script>
