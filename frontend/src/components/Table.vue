<template>
  <section>
    <b-table
      :data="data"
      :columns="columns"
      :default-sort="sort"
      paginated
      hoverable
      :row-class="(row, index) => rowClass"
      @click="(row, index) => $emit('row-click', row, index)"
    >
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
  data() {
    return {
      rowClass: {
        "has-row-style": true,
      },
    };
  },
});
</script>

<style lang="scss">
.has-row-style:hover {
  cursor: pointer;
}
</style>
