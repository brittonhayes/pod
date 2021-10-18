<template>
  <div class="modal-card" style="width: auto">
    <div class="modal-card-head">
      New Project
    </div>
    <section class="modal-card-body">
      <b-field label="Project">
        <b-input v-model="form.name" icon="tasks" required />
      </b-field>
      <b-field label="Client">
        <b-autocomplete
          v-model="form.client"
          :data="clients"
          field="name"
          placeholder="Select a client"
          required
          @select="(option) => (selected = option)"
        >
          <template #empty>No results found</template>
        </b-autocomplete>
      </b-field>
      <b-field label="Summary">
        <b-input v-model="form.summary" type="textarea" maxlength="250" />
      </b-field>
    </section>
    <footer class="modal-card-foot">
      <b-button label="Go back" @click="toggleModal" />
      <b-button label="Save" @click="SubmitProject" type="is-success" />
    </footer>
  </div>
</template>

<script lang="ts">
import { SAVE, REFRESH, TOGGLE, Namespace } from "@/store/mutations";

import Vue from "vue";
import { mapMutations, mapState } from "vuex";

export default Vue.extend({
  props: {
    name: String,
    summary: String,
    client: String,
  },
  data() {
    return {
      form: {
        name: this.name,
        summary: this.summary,
        client: this.client,
      },
    };
  },
  methods: {
    ...mapMutations({
      toggleModal: Namespace.Projects + TOGGLE,
    }),
    SubmitProject: function() {
      this.$store.commit(Namespace.Projects + SAVE, this.form);
      this.$store.commit(Namespace.Projects + REFRESH);
      this.$store.commit(Namespace.Projects + TOGGLE);
    },
  },
  computed: {
    ...mapState({
      clients: (state: any) => state.clients.clients,
    }),
  },
});
</script>
