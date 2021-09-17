import Vue from "vue";
export default Vue.extend({
  computed: {
    title: function() {
      //@ts-ignore
      return this.$router.currentRoute.meta.title;
    },
    subtitle: function() {
      //@ts-ignore
      return this.$router.currentRoute.meta.subtitle;
    },
  },
});
