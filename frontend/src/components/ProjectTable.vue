<template>
  <section class="table-container">
    <b-table
      :data="projects"
      :columns="columns"
      :default-sort="sort"
      paginated
      detailed
    >
      <template #detail="props">
        <small>{{ props.row.summary }}</small>
      </template>
      <template #empty>
        <div class="notification has-text-centered">No projects</div>
      </template>
    </b-table>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { ProjectColumns, ProjectSort } from "@/lib/projects";

export default Vue.extend({
  data() {
    return {
      columns: ProjectColumns,
      sort: ProjectSort,
    };
  },
  props: {
    projects: {
      type: Array,
      required: true,
    },
  },
});
</script>
