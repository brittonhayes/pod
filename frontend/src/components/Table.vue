<template>
  <section>
    <b-table :data="data" :columns="columns" :default-sort="sort" paginated>
      <template #empty>
        <div class="notification has-text-centered">No results</div>
      </template>
    </b-table>
  </section>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  props: {
    data: {
      type: Array,
      required: true,
    },
    columns: {
      type: Array,
      required: true,
    },
    sort: {
      type: Array,
      required: true,
    },
  },
});
</script>
