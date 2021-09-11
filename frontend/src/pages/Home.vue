<template>
  <div>
    <Banner :title="title" :subtitle="subtitle" />
    <div class="columns">
      <div class="column my-5 mx-3 box has-background-success has-text-white">
        Recent projects
      </div>
      <div class="column my-5 mx-3 box">
        Clients
      </div>
      <div class="column my-5 mx-3 box">
        Sounds
      </div>
      <div class="column my-5 mx-3 box">
        Settings
      </div>
    </div>
    <Table />
  </div>
</template>

<script>
import Banner from "@/components/Banner.vue";
import Table from "@/components/Table.vue";
export default {
  components: {
    Banner,
    Table,
  },
  data() {
    return {
      title: "Pod 🌱",
      subtitle: "The audio professional's productivity center",
    };
  },
};
</script>
