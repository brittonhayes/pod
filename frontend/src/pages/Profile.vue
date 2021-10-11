<template>
  <container>
    <Contact />
    <b-field class="my-5">
      <p class="control buttons">
        <b-button
          icon-left="chevron-left"
          label="Go back"
          type="is-light"
          tag="router-link"
          to="/clients"
          expanded
        ></b-button>
      </p>
    </b-field>
  </container>
</template>

<script lang="ts">
import Vue from "vue";
import Container from "@/components/Container.vue";
import Contact from "@/components/Contact.vue";

export default Vue.extend({
  name: "Profile",
  components: {
    Container,
    Contact,
  },
  methods: {},
});
</script>
