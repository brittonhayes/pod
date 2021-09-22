<template>
  <section>
    <div class="mb-6">
      <nav class="level">
        <div class="level-item level-left">
          <figure class="image is-96x96 mr-5">
            <img
              class="is-rounded"
              src="https://bulma.io/images/placeholders/96x96.png"
            />
          </figure>
          <div class="block">
            <h1 class="title is-size-2">
              Client Name
            </h1>
            <small>
              Last updated: <em>{{ new Date().toLocaleString() }}</em>
            </small>
          </div>
        </div>
      </nav>
    </div>

    <b-field label="Summary">
      <b-input type="textarea"></b-input>
    </b-field>

    <b-field class="mb-5">
      <p class="control buttons">
        <b-button label="Edit"></b-button>
        <b-button label="Save" type="is-primary" disabled></b-button>
      </p>
    </b-field>

    <h2 class="is-size-6 has-text-weight-bold pb-3">Actions</h2>
    <div class="buttons">
      <b-button label="View projects"></b-button>
      <b-button label="View clips"></b-button>
      <b-button label="View folder"></b-button>
    </div>
    <h2 class="is-size-6 has-text-weight-bold pb-3">Recent activity</h2>
    <div class="block">
      <b-skeleton class="py-3" :height="35"></b-skeleton>
      <b-skeleton class="py-3" :height="35"></b-skeleton>
      <b-skeleton class="py-3" :height="35"></b-skeleton>
      <b-skeleton class="py-3" :height="35"></b-skeleton>
    </div>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import Page from "@/mixins/Page";

export default Vue.extend({
  mixins: [Page],
});
</script>
