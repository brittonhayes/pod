<template>
  <div>
    <banner :title="title" :subtitle="subtitle" />
    <section>
      <div class="block">
        <b-button
          label="New Project"
          icon-left="plus"
          @click="isOpen = !isOpen"
        />
      </div>
      <b-modal v-model="isOpen" has-modal-card :width="1500">
        <project-form v-bind="project"></project-form>
      </b-modal>
      <Cards :items="projects" color="light" />
    </section>
  </div>
</template>

<script>
import Page from "@/mixins/Page.js";
import Cards from "@/components/Cards.vue";
import Banner from "@/components/Banner.vue";
import ProjectForm from "@/components/ProjectForm.vue";

export default {
  data() {
    return {
      isOpen: false,
      projects: [],
      classes: {
        collapse: {
          panel: true,
        },
      },
      project: {
        name: "",
        summary: "",
        client: "",
      },
    };
  },
  mixins: [Page],
  components: {
    Banner,
    Cards,
    ProjectForm,
  },
  methods: {
    LoadProjects: function() {
      window.backend.Storage.ListProjects()
        .then((res) => {
          this.projects = res;
          return res;
        })
        .catch((err) => {
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
};
</script>
