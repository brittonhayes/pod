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

<script>
import NavButton from "@/components/NavButton.vue";
export default {
  components: { NavButton },
  data() {
    return {
      routes: this.$router.options.routes,
    };
  },
  compontents: {
    NavButton,
  },
  computed: {
    activeRoute: function() {
      console.log(this.$store.getters.getActiveRoute);
      return this.$store.getters.getActiveRoute;
    },
  },
};
</script>

<style>
.logo-light {
  opacity: 0.9;
}
</style>
