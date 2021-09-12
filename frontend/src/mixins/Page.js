export default {
  computed: {
    title: function() {
      return this.$router.currentRoute.meta.title;
    },
    subtitle: function() {
      return this.$router.currentRoute.meta.subtitle;
    },
  },
};
