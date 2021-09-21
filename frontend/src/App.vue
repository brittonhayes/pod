<template>
  <div id="app">
    <sidebar />
    <router-view class="margin-body"></router-view>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Sidebar from "./components/Sidebar.vue";
export default Vue.extend({
  name: "app",
  components: { Sidebar },
});
</script>

<style lang="scss">
.margin-body {
  margin-top: 20px;
  margin-bottom: 20px;
  margin-left: 120px;
  margin-right: 30px;
}
</style>
