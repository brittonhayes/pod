<template>
  <container>
    <banner :title="title" :subtitle="subtitle">
      <template slot="buttons">
        <div class="block buttons">
          <b-button
            icon-left="plus"
            label="Create project"
            type="is-success is-light"
            @click="isOpen = !isOpen"
          />
          <b-button icon-left="redo" type="is-light" @click="LoadProjects" />
        </div>
      </template>
    </banner>
    <section>
      <b-modal v-model="isOpen" has-modal-card :width="1500">
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
