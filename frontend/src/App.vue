<template>
  <div id="app">
    <sidebar />
    <router-view class="margin-body"></router-view>
  </div>
</template>

<script>
import Sidebar from "./components/Sidebar.vue";
export default {
  name: "app",
  components: { Sidebar },
};
</script>

<style lang="scss">
.margin-body {
  margin-top: 20px;
  margin-bottom: 20px;
  margin-left: 120px;
  margin-right: 30px;
}
</style>
