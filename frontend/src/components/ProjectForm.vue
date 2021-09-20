<template>
  <div class="modal-card" style="width: auto">
    <div class="modal-card-head">
      New Project
    </div>
    <section class="modal-card-body">
      <b-field label="Name">
        <b-input v-model="name" />
      </b-field>
      <b-field label="Client">
        <b-input v-model="client" maxlength="30" />
      </b-field>
      <b-field label="Summary">
        <b-input v-model="summary" type="textarea" maxlength="250" />
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
  methods: {
    SubmitProject: function() {
      let p = JSON.stringify({
        name: this.name,
        summary: this.summary,
        client: this.client,
      });

      //@ts-ignore
      window.backend.Storage.SaveProject(p)
        .then((res) => {
          console.debug(res);
          this.$buefy.snackbar.open({
            message: `Submitted ${this.name}`,
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
      this.LoadProjects();
    },
  },
});
</script>
