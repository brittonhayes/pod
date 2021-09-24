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

import { ClientSort, ClientColumns } from "@/types/client";

import { mapMutations, mapState } from "vuex";
import { UPDATE_FROM_DB } from "@/store/modules/mutations";

export default Vue.extend({
  mixins: [Page],
  components: {
    Container,
    Banner,
    Table,
  },
  data() {
    return {
      sort: ClientSort,
      columns: ClientColumns,
    };
  },
  methods: {
    ...mapMutations({ updateFromDB: UPDATE_FROM_DB }),
  },
  computed: {
    ...mapState({
      items: (state: any) => state.clients.clients,
    }),
  },
});
</script>
