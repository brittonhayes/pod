<template>
  <section>
    <b-sidebar
      type="is-success"
      :reduce="true"
      :fullheight="true"
      :overlay="false"
      open
    >
      <nav-button
        v-for="route in routes"
        :key="route.path"
        :icon="route.meta.icon"
        :size="route.meta.iconSize"
        :to="route.path"
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
};
</script>

<style>
.logo-light {
  opacity: 0.9;
}
</style>
