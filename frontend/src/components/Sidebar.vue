<template>
  <section>
    <b-sidebar
      type="is-success"
      :reduce="true"
      :fullheight="true"
      :overlay="false"
      open
    >
      <b-icon
        size="is-large"
        icon="leaf"
        type="is-light"
        class="m-2 pl-2 logo-light"
      ></b-icon>
      <NavButton icon="tasks" />
      <NavButton icon="address-book" />
      <NavButton icon="volume-up" />
      <NavButton icon="cog" />
    </b-sidebar>
  </section>
</template>

<script>
import NavButton from "@/components/NavButton.vue";
export default {
  components: { NavButton },
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
