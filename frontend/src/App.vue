<template>
  <div id="app">
    <Sidebar />
    <router-view class="margin-body"></router-view>
  </div>
</template>

<script>
import Sidebar from "@/components/Sidebar.vue";
export default {
  name: "app",
  components: { Sidebar },
};
</script>

<style>
.margin-body {
  margin-top: 20px;
  margin-left: 120px;
}
</style>
