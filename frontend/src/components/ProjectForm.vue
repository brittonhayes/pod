<template>
  <div class="modal-card" style="width: auto">
    <div class="modal-card-head">
      New Project
    </div>
    <section class="modal-card-body">
      <b-field label="Name">
        <b-input v-model="internalName" />
      </b-field>
      <b-field label="Client">
        <b-input v-model="internalClient" maxlength="30" />
      </b-field>
      <b-field label="Summary">
        <b-input v-model="internalSummary" type="textarea" maxlength="250" />
      </b-field>
    </section>
    <footer class="modal-card-foot">
      <b-button label="Go back" @click="$parent.close()" />
      <b-button label="Save" @click="SubmitProject" type="is-success" />
    </footer>
  </div>
</template>

<script>
import Vue from "vue";
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
      internalName: this.name,
      internalSummary: this.summary,
      internalClient: this.client,
    };
  },
  methods: {
    SubmitProject: function() {
      const project = {
        Name: this.internalName,
        Summary: this.internalSummary,
        Client: this.internalClient,
      };

      window.backend.Storage.SaveProject(project)
        .then(() => {
          this.$buefy.snackbar.open({
            message: `Submitted ${this.internalName}`,
            type: "is-success",
          });
        })
        .catch((err) => {
          this.$buefy.snackbar.open({
            message: `Failed to submit - ${err}`,
            type: "is-danger",
          });
          console.error(err);
        });

      this.$parent.close();
    },
  },
});
</script>
