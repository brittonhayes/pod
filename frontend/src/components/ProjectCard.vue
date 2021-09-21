<template>
  <section class="box has-border-project">
    <article class="content">
      <h2 class="title is-size-5 has-text-project">
        {{ project }}
      </h2>
      <p class="subtitle is-size-6 has-text-client">{{ client }}</p>
      <p class="is-hidden-touch">{{ content }}</p>
    </article>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  name: "ProjectCard",
  props: {
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
.has-text-project {
  color: #161616;
}
.has-text-client {
  color: #aaaaaa;
}

.has-border-project {
  border-style: solid;
  border-width: 1px;
  border-color: #e6e6e6;
}
</style>
