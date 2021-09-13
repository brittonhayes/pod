<template>
  <div>
    <div class="columns is-multiline">
      <div
        v-for="(item, index) in paginatedItems"
        :key="index"
        class="column is-one-third"
      >
        <card :title="item.name" />
      </div>
    </div>
    <section class="section">
      <b-pagination
        :total="total"
        :current.sync="current"
        :per-page="perPage"
        color="light"
        simple
      ></b-pagination>
    </section>
  </div>
</template>

<script>
import Color from "@/mixins/Color.js";
import Card from "./Card.vue";
export default {
  components: { Card },
  name: "Cards",
  mixins: [Color],
  data() {
    return {
      classes: {
        card: {
          card: true,
        },
        cardHeader: {
          "is-size-6": true,
          "pb-5": true,
        },
        cardContent: {
          content: true,
          "is-size-7": true,
        },
      },
    };
  },
  props: {
    current: {
      type: Number,
      default: () => {
        return 1;
      },
    },
    perPage: {
      type: Number,
      default: () => {
        return 6;
      },
    },
    items: {
      type: Array,
      default: () => [
        {
          id: 1,
          title: "Card",
          body: "Card body content",
        },
      ],
    },
  },
  computed: {
    total() {
      return this.items.length;
    },
    paginatedItems() {
      let page_number = this.current - 1;

      return this.items.slice(
        page_number * this.perPage,
        (page_number + 1) * this.perPage
      );
    },
  },
};
</script>
