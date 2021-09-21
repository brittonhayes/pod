<template>
  <div class="container">
    <section class="section is-small">
      <div class="columns is-mobile is-centered">
        <div class="column is-two-thirds">
          <slot />
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  name: "Container",
});
</script>

<style lang="scss">
.center-container {
  display: flex;
  width: 100%;
  flex-grow: 1;
  justify-content: center;
  // @include desktop-only {
  margin-top: 40px;
  // }
}
</style>
