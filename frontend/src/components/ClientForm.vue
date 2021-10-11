<template>
  <div class="modal-card" style="width: auto">
    <div class="modal-card-head">
      New Client
    </div>
    <section class="modal-card-body">
      <b-field label="Name">
        <b-input v-model="form.name" required icon="user" />
      </b-field>
      <b-field label="Email">
        <b-input v-model="form.email" required icon="envelope" />
      </b-field>
      <b-field label="Phone">
        <b-input v-model="form.phone" required icon="phone" />
      </b-field>
    </section>
    <footer class="modal-card-foot">
      <b-button label="Go back" @click="toggleModal" />
      <b-button label="Save" @click="SubmitClient" type="is-success" />
    </footer>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import {
  SUBMIT_FORM,
  UPDATE_FROM_DB,
  TOGGLE_ENABLED,
  Namespace,
} from "@/store/mutations";
import { mapMutations } from "vuex";

export default Vue.extend({
  props: {
    name: String,
    description: String,
    email: String,
    phone: String,
  },
  data() {
    return {
      form: {
        name: this.name,
        email: this.email,
        phone: this.phone,
      },
    };
  },
  methods: {
    ...mapMutations({
      toggleModal: Namespace.Clients + TOGGLE_ENABLED,
    }),
    SubmitClient: function() {
      this.$store.commit(Namespace.Clients + SUBMIT_FORM, this.form);
      this.$store.commit(Namespace.Clients + UPDATE_FROM_DB);
      this.$store.commit(Namespace.Clients + TOGGLE_ENABLED);
    },
  },
});
</script>
