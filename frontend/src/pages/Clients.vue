<template>
  <container>
    <Banner :title="title" :subtitle="subtitle">
      <template slot="buttons">
        <div class="block buttons is-hidden-mobile">
          <b-button
            icon-left="plus"
            label="Create client"
            type="is-success is-light"
            @click="toggleModal"
          />
          <b-button icon-left="redo" type="is-light" @click="updateFromDB" />
        </div>
      </template>
    </Banner>
    <section>
      <b-modal
        :active="modalEnabled"
        @close="toggleModal"
        :width="1500"
        can-cancel
        has-modal-card
      >
        <client-form v-bind="client"></client-form>
      </b-modal>
    </section>
    <Table
      :data="items"
      :columns="columns"
      :sort="sort"
      @row-click="(row) => viewClient(row)"
    />
  </container>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import Container from "@/components/Container.vue";
import Table from "@/components/Table.vue";
import Banner from "@/components/Banner.vue";
import Page from "@/mixins/Page";

import ClientForm from "@/components/ClientForm.vue";
import { ClientSort, ClientColumns, Client } from "@/types/client";

import { mapMutations, mapState, mapGetters } from "vuex";
import {
  IS_ENABLED,
  UPDATE_FROM_DB,
  TOGGLE_ENABLED,
  Namespace,
} from "@/store/mutations";

export default Vue.extend({
  mixins: [Page],
  components: {
    Container,
    Banner,
    Table,
    ClientForm,
  },
  data() {
    return {
      client: {
        type: Object as PropType<Client>,
        default: () => {},
      },
      sort: ClientSort,
      columns: ClientColumns,
    };
  },
  mounted() {
    this.$store.commit(Namespace.Clients + UPDATE_FROM_DB);
  },
  methods: {
    ...mapMutations({
      updateFromDB: Namespace.Clients + UPDATE_FROM_DB,
      toggleModal: Namespace.Clients + TOGGLE_ENABLED,
    }),
    viewClient: function(row: Client) {
      console.log(row);
      this.$router.push("/profile/" + row.id);
    },
  },
  computed: {
    ...mapGetters({ modalEnabled: Namespace.Clients + IS_ENABLED }),
    ...mapState({
      items: (state: any) => state.clients.clients,
    }),
  },
});
</script>
