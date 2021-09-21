import VueRouter, { Route } from "vue-router";

declare module "vue/types/vue" {
  import Vue from "vue";
  interface Vue {
    $router: VueRouter;
  }
  export default Vue;
}
