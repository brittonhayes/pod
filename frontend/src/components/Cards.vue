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
      ></b-pagination>
    </section>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Card } from "@/types/cards";
import CardComponent from "@/components/Card.vue";
export default Vue.extend({
  name: "Cards",
  components: { card: CardComponent },
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
    total(): number {
      let list = this.items as Array<Card>;
      return list.length;
    },
    paginatedItems(): Array<Card> {
      let page_number = this.current - 1;

      return this.items.slice(
        page_number * this.perPage,
        (page_number + 1) * this.perPage
      ) as Array<Card>;
    },
  },
});
</script>
