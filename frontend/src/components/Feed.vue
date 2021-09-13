<template>
  <article class="media my-3">
    <div class="media-content">
      <div class="content">
        <slot v-if="!empty" name="header" />
        <slot v-if="!empty" name="subtitle" />
        <p v-if="empty">
          <b-skeleton :width="150 * i" :animated="false" active></b-skeleton>
          <b-skeleton :animated="false" height="80px"></b-skeleton>
        </p>
      </div>
      <nav class="level is-mobile">
        <div class="level-left">
          <a class="level-item">
            <slot v-if="!empty" name="icon-left" />
            <span v-if="empty" class="icon is-small">
              <b-skeleton :animated="false"></b-skeleton>
            </span>
          </a>
          <a class="level-item">
            <slot v-if="!empty" name="icon-right" />
            <span v-if="empty" class="icon is-small">
              <b-skeleton :animated="false"></b-skeleton>
            </span>
          </a>
        </div>
      </nav>
    </div>
  </article>
</template>

<script>
export default {
  data() {
    return {
      empty: false,
    };
  },
};
</script>
