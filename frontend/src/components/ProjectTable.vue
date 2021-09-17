<template>
  <section class="table-container">
    <b-table
      :data="data"
      :columns="columns"
      :default-sort="sort"
      scrollable
      hoverable
      paginated
    >
      <template #empty>
        <div class="notification has-text-centered">No projects</div>
      </template>
    </b-table>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { DefaultProjects, ProjectColumns, ProjectSort } from "../lib/projects";

export default Vue.extend({
  data() {
    return {
      data: DefaultProjects,
      columns: ProjectColumns,
      sort: ProjectSort,
    };
  },
});
</script>
