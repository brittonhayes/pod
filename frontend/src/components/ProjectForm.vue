<template>
  <div class="modal-card" style="width: auto">
    <div class="modal-card-head">
      New Project
    </div>
    <section class="modal-card-body">
      <b-field label="Name">
        <b-input v-model="form.Name" />
      </b-field>
      <b-field label="Client">
        <b-input v-model="form.Client" maxlength="30" />
      </b-field>
      <b-field label="Summary">
        <b-input v-model="form.Summary" type="textarea" maxlength="250" />
      </b-field>
    </section>
    <footer class="modal-card-foot">
      <b-button label="Go back" @click="toggleModal" />
      <b-button label="Save" @click="SubmitProject" type="is-success" />
    </footer>
  </div>
</template>

<script lang="ts">
import {
  SUBMIT_FORM,
  UPDATE_FROM_DB,
  TOGGLE_ENABLED,
} from "@/store/modules/mutations";
import Vue from "vue";
import { mapMutations } from "vuex";
export default Vue.extend({
  props: {
    name: {
      type: String,
    },
    summary: {
      type: String,
    },
    client: {
      type: String,
    },
  },
  data() {
    return {
      form: {
        Name: this.name,
        Summary: this.summary,
        Client: this.client,
      },
    };
  },
  methods: {
    ...mapMutations({
      toggleModal: TOGGLE_ENABLED,
    }),
    SubmitProject: function() {
      this.$store.commit(SUBMIT_FORM, this.form);
      this.$store.commit(UPDATE_FROM_DB);
      this.$store.commit(TOGGLE_ENABLED);
    },
  },
});
</script>
