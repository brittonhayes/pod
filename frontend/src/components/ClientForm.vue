<template>
  <div class="modal-card" style="width: auto">
    <div class="modal-card-head">
      New Client
    </div>
    <section class="modal-card-body">
      <b-field label="Name">
        <b-input v-model="form.Name" required icon="user" />
      </b-field>
      <b-field label="Email">
        <b-input v-model="form.Email" required icon="envelope" />
      </b-field>
      <b-field label="Phone">
        <b-input v-model="form.Phone" required icon="phone" />
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
  Mutator,
  CLIENTS,
} from "@/store/mutations";
import { mapMutations } from "vuex";

const mu = new Mutator(CLIENTS);

// name: string;
// description?: string;
// email?: string;
// phone?: string;
// social?: Object;

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
      },
    };
  },
  methods: {
    ...mapMutations({
      toggleModal: mu.Mutation(TOGGLE_ENABLED),
    }),
    SubmitClient: function() {
      this.$store.commit(mu.Mutation(SUBMIT_FORM), this.form);
      this.$store.commit(mu.Mutation(UPDATE_FROM_DB));
      this.$store.commit(mu.Mutation(TOGGLE_ENABLED));
    },
  },
});
</script>
