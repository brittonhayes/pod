<template>
  <b-table :data="data" :columns="columns"></b-table>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      data: [
        {
          id: 1,
          name: "Mythbusters",
          client: "Discovery",
          type: "VO",
          date: "2016-10-15 13:43:27",
          link: "🔗",
        },
        {
          id: 2,
          name: "Spider-Man - Miles Morales",
          client: "Marvel Studios",
          type: "VO Mixing",
          date: "2016-12-15 06:00:53",
          link: "🔗",
        },
        {
          id: 3,
          name: "My Podcast",
          client: "Me",
          type: "Podcast",
          date: "2016-04-26 06:26:28",
          link: "🔗",
        },
        {
          id: 4,
          name: "Game of Thrones",
          client: "HBO",
          type: "Foley",
          date: "2016-04-26 06:26:28",
          link: "🔗",
        },
        {
          id: 5,
          name: "Untitled Album",
          client: "James",
          type: "Production",
          date: "2016-04-26 06:26:28",
          link: "🔗",
        },
      ],
      columns: [
        {
          field: "id",
          label: "ID",
          width: "40",
          numeric: true,
        },
        {
          field: "name",
          label: "Project",
        },
        {
          field: "client",
          label: "Client",
        },
        {
          field: "type",
          label: "Type",
        },
        {
          field: "date",
          label: "Date",
        },
        {
          field: "link",
          label: "Link",
        },
      ],
    };
  },
});
</script>
