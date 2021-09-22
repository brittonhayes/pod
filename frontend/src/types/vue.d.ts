import VueRouter, { Route } from "vue-router";
import Vue, { VNode } from "vue";
import { Store } from "vuex";

declare module "vue/types/vue" {
  import Vue from "vue";
  interface Vue {
    $router: VueRouter;
  }
  export default Vue;
}

declare module "@vue/runtime-core" {
  // declare your own store states
  interface State {
    count: number;
  }

  // provide typings for `this.$store`
  interface ComponentCustomProperties {
    $store: Store<State>;
  }
}

declare global {
  namespace JSX {
    // tslint:disable no-empty-interface
    interface Element extends VNode {}
    // tslint:disable no-empty-interface
    interface ElementClass extends Vue {}
    interface IntrinsicElements {
      [elem: string]: any;
    }
  }
}
