<template>
  <section>
    <div class="mb-6">
      <nav class="level">
        <div class="level-item level-left">
          <figure class="image is-96x96 mr-5">
            <b-image
              :src="
                `https://avatars.dicebear.com/api/initials/${client.name}.svg`
              "
              ratio="96by96"
              rounded
            ></b-image>
          </figure>
          <div class="block">
            <h1 class="title is-size-2">
              {{ client.name }}
              <span class="has-text-danger" v-show="edit">*</span>
            </h1>
            <small>
              Last updated:
              <em>{{ updatedAt }}</em>
            </small>
          </div>
        </div>
        <div class="level-item level-right"></div>
      </nav>
    </div>

    <section class="mb-4">
      <b-field label="Name">
        <b-input
          type="input"
          :disabled="!edit"
          lazy
          v-model="client.name"
        ></b-input>
      </b-field>
      <b-field label="Email">
        <b-input
          type="input"
          :disabled="!edit"
          lazy
          v-model="client.email"
        ></b-input>
      </b-field>
      <b-field label="Phone">
        <b-input
          type="tel"
          :disabled="!edit"
          lazy
          v-model="client.phone"
        ></b-input>
      </b-field>
      <b-field label="Summary">
        <b-input
          :disabled="!edit"
          type="textarea"
          lazy
          v-model="client.description"
          ex
        ></b-input>
      </b-field>
    </section>

    <b-field class="mb-5 buttons">
      <b-field>
        <b-button
          outlined
          expanded
          icon-left="pencil-alt"
          label="Edit"
          v-show="!edit"
          @click="toggleEdit"
        ></b-button>
      </b-field>
      <b-field grouped>
        <b-field expanded>
          <b-button
            label="Undo"
            type="is-light"
            v-show="edit"
            :disabled="!edit"
            @click="toggleEdit"
          ></b-button>

          <b-button
            outlined
            label="Save"
            icon-left="save"
            type="is-primary"
            v-show="edit"
            :disabled="!edit"
            @click="updateClient"
          ></b-button>
        </b-field>
      </b-field>
    </b-field>
    <h2 class="is-size-6 has-text-weight-bold pb-3">Actions</h2>
    <div class="buttons">
      <b-button label="View projects"></b-button>
      <b-button label="View clips"></b-button>
      <b-button label="View folder"></b-button>
    </div>
    <h2 class="is-size-6 has-text-weight-bold pb-3">Recent activity</h2>
    <div class="block">
      <b-skeleton class="py-3" :animated="false" :height="35"></b-skeleton>
      <b-skeleton class="py-3" :animated="false" :height="35"></b-skeleton>
      <b-skeleton class="py-3" :animated="false" :height="35"></b-skeleton>
      <b-skeleton class="py-3" :animated="false" :height="35"></b-skeleton>
    </div>
  </section>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import Page from "@/mixins/Page";
import { Client } from "@/types/client";
import { Namespace, UPDATE } from "@/store/mutations";

export default Vue.extend({
  mixins: [Page],
  data() {
    return {
      edit: false,
      client: {
        type: Object as PropType<Client>,
        default: () => {},
      },
    };
  },
  beforeMount() {
    this.loadClient();
  },
  methods: {
    updateClient: function() {
      this.$store.commit(Namespace.Clients + UPDATE, this.client);
      this.toggleEdit();
      this.loadClient();
    },
    loadClient: function() {
      let id = parseInt(this.$route.params.id);
      window.backend.Storage.FindClient(id)
        .then((c) => {
          this.client = c;
        })
        .catch((e) => {
          console.error(e);
        });
    },
    toggleEdit: function() {
      this.edit = !this.edit;
      if (!this.edit) {
        this.loadClient();
      }
    },
  },
  computed: {
    updatedAt: function(): string {
      //@ts-ignore
      return new Date(this.client.updated_at).toLocaleString();
    },
  },
});
</script>
