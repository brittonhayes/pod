<template>
  <div class="card my-5">
    <article class="card-content">
      <div class="media">
        <div class="media-left">
          <figure class="image is-48x48">
            <b-icon icon="tasks" />
          </figure>
        </div>
        <div class="media-content">
          <p class="title is-6" v-text="project" />
          <p class="subtitle is-7 has-text-danger" v-text="client" />
        </div>
      </div>
    </article>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  name: "ProjectCard",
  props: {
    detailed: {
      type: Boolean,
      default: () => {
        return false;
      },
    },
    project: {
      type: String,
      required: true,
      default: () => {
        return "Project Name";
      },
    },
    client: {
      type: String,
      required: true,
      default: () => {
        return "Client Name";
      },
    },
    content: {
      type: String,
      default: () => {
        return `Lorem ipsum dolor sit, amet consectetur adipisicing elit. Dicta,
        reprehenderit suscipit saepe explicabo accusantium harum a? Ab modi eum,
        quos natus ullam cumque deleniti officiis accusamus temporibus molestiae
        quia hic.`;
      },
    },
  },
});
</script>

<style lang="scss" scoped>
.has-bg-project {
}

.has-text-project {
  color: #161616;
}
.has-text-client {
  color: #aaaaaa;
  font-family: monospace;
}

.has-border-project {
  border-style: solid;
  border-width: 1px;
  border-color: #e6e6e6;
}
</style>
