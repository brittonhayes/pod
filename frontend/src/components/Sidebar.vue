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
import { SET_ACTIVE, Mutator, ROUTER } from "@/store/mutations";
import { mapState, mapMutations } from "vuex";
import { Route } from "vue-router";

const mu = new Mutator(ROUTER);

export default Vue.extend({
  watch: {
    $route(to: Route) {
      this.$store.commit(mu.Mutation(SET_ACTIVE), to.path);
    },
  },
  components: { NavButton },
  methods: {
    ...mapMutations({
      setActive: mu.Mutation(SET_ACTIVE),
    }),
  },
  computed: {
    routes: function() {
      let r = this.$router.options.routes;
      r = r?.filter(function(item) {
        return !item.path.includes("/:");
      });

      return r;
    },
    ...mapState({
      activeRoute: (state: any) => state.router.active,
    }),
  },
});
</script>

<style>
.logo-light {
  opacity: 0.9;
}
</style>
