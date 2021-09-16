<template>
  <div>
    <banner :title="title" :subtitle="subtitle" />
    <section>
      <div class="block pl-5 buttons">
        <b-button label="New project" @click="isOpen = !isOpen" />
        <b-button icon-left="redo" type="is-light" @click="LoadProjects" />
      </div>
      <b-field class="px-5 my-5">
        <b-input
          type="search"
          icon="search"
          placeholder="Search..."
          v-model="query"
          expanded
        ></b-input>
      </b-field>

      <b-modal v-model="isOpen" has-modal-card :width="1500">
        <project-form v-bind="project"></project-form>
      </b-modal>
      <section
        v-for="project in projects"
        :key="project.id"
        class="notification is-light mx-5"
      >
        <article class="media">
          <div class="media-content">
            <div class="content">
              <strong>{{ project.name }}</strong>
              <br />
              <small>{{ project.client }}</small>
              <br />
            </div>
          </div>
          <figure class="media-right">
            <b-button
              icon-left="angle-right"
              size="is-large"
              type="is-text"
            ></b-button>
          </figure>
        </article>
      </section>
      <!-- <Cards :items="projects" color="light" /> -->
    </section>
  </div>
</template>

<script>
import Page from "@/mixins/Page.js";
// import Cards from "@/components/Cards.vue";
import Banner from "@/components/Banner.vue";
import ProjectForm from "@/components/ProjectForm.vue";

export default {
  mixins: [Page],
  components: {
    Banner,
    // Cards,
    ProjectForm,
  },
  data() {
    return {
      isOpen: false,
      query: "",
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
