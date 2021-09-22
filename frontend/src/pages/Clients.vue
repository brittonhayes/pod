<template>
  <container>
    <Banner :title="title" :subtitle="subtitle" />
    <Table :data="items" :columns="columns" :sort="sort" />
  </container>
</template>

<script lang="ts">
import Vue from "vue";
import Container from "@/components/Container.vue";
import Table from "@/components/Table.vue";
import Banner from "@/components/Banner.vue";
import Page from "@/mixins/Page";

import { DefaultClients, ClientSort, ClientColumns } from "@/types/client";

export default Vue.extend({
  mixins: [Page],
  components: {
    Container,
    Banner,
    Table,
  },
  data() {
    return {
      items: DefaultClients,
      sort: ClientSort,
      columns: ClientColumns,
    };
  },
});
</script>
