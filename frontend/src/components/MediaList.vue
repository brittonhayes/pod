<template>
  <section class="notification is-light mx-5">
    <article class="media">
      <div class="media-content">
        <div class="content">
          <strong>{{ content.title }}</strong>
          <br />
          <small>{{ content.subtitle }}</small>
          <br />
        </div>
      </div>
      <figure class="media-right">
        <b-button
          icon-left="angle-right"
          size="is-large"
          type="is-text"
        ></b-button>
      </figure>
    </article>
  </section>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";

interface Media {
  title: string;
  subtitle?: string;
}

export default Vue.extend({
  props: {
    content: {
      type: Object as PropType<Media>,
      required: true,
      validator(content: Media): boolean {
        return !!content.title;
      },
    },
  },
});
</script>
