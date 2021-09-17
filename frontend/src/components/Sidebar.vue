<template>
  <section>
    <b-sidebar
      type="is-success"
      :can-cancel="false"
      :reduce="true"
      :fullheight="true"
      :overlay="false"
      :open="true"
    >
      <nav-button
        v-for="route in routes"
        :ref="route.path"
        :key="route.path"
        :icon="route.meta.icon"
        :size="route.meta.iconSize"
        :to="route.path"
        :active="route.path == activeRoute"
        color="success"
        >{{ route.path }}</nav-button
      >
    </b-sidebar>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import NavButton from "./NavButton.vue";

export default Vue.extend({
  components: { NavButton },
  computed: {
    routes: function() {
      return this.$router.options.routes;
    },
    activeRoute: function() {
      return this.$store.getters.getActiveRoute;
    },
  },
});
</script>

<style>
.logo-light {
  opacity: 0.9;
}
</style>
