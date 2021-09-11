<template>
  <section class="hero is-light">
    <div class="hero-body">
      <h1 class="title is-size-3 has-text-success has-text-weight-light pb-2">
        {{ title }}
      </h1>
      <h2 class="subtitle is-size-6">
        {{ subtitle }}
      </h2>
    </div>
  </section>
</template>

<script>
export default {
  name: "Banner",
  props: {
    title: {
      type: String,
      required: true,
    },
    subtitle: {
      type: String,
    },
  },
  data() {
    return {};
  },
};
</script>
