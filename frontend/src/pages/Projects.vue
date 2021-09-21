<template>
  <div>
    <banner :title="title" :subtitle="subtitle">
      <div class="block buttons">
        <b-button
          icon-left="plus"
          label="Create project"
          type="is-success is-light"
          @click="isOpen = !isOpen"
        />
        <b-button icon-left="redo" type="is-light" @click="LoadProjects" />
      </div>
    </banner>
    <section>
      <project-table :projects="projects" class="px-4"></project-table>
      <b-modal v-model="isOpen" has-modal-card :width="1500">
        <project-form v-bind="project"></project-form>
      </b-modal>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "@/mixins/Page";
import Banner from "@/components/Banner.vue";
import ProjectTable from "@/components/ProjectTable.vue";
import ProjectForm from "@/components/ProjectForm.vue";

export default Vue.extend({
  mixins: [Page],
  components: {
    Banner,
    ProjectTable,
    ProjectForm,
  },
  data() {
    return {
      isOpen: false,
      projects: [],
      project: {
        name: "",
        summary: "",
        client: "",
      },
    };
  },
  methods: {
    LoadProjects: function() {
      const that = this;
      //@ts-ignore
      window.backend.Storage.ListProjects()
        .then((res: any) => {
          that.projects = res;
          return res;
        })
        .catch((err: Error) => {
          console.error(err);
          return [{}];
        });
    },
  },
  mounted() {
    this.LoadProjects();
  },
  computed: {
    classObject: function() {
      return {
        panel: true,
      };
    },
  },
});
</script>
