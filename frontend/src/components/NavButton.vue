<template>
  <b-button
    :to="route"
    size="is-medium"
    class="m-2"
    :type="color"
    :icon-left="icon"
  />
</template>

<script>
export default {
  data() {
    return {
      margin: "m-2",
    };
  },
  props: {
    route: {
      type: String,
    },
    color: {
      type: String,
      default: "is-success",
      validator: function(value) {
        return (
          [
            "is-primary",
            "is-info",
            "is-success",
            "is-warning",
            "is-danger",
            "is-dark",
            "is-light",
          ].indexOf(value) !== -1
        );
      },
    },
    icon: {
      type: String,
      default: "user",
    },
  },
};
</script>
