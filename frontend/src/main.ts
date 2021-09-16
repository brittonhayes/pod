import "core-js/stable";
import "regenerator-runtime/runtime";
import * as Wails from "@wailsapp/runtime";
import Vue from "vue";
import "buefy/dist/buefy.css";
import Buefy from "buefy";

import FontAwesomeIcon from "./icons";
import App from "./App.vue";
import router from "./router/router";
import store from "./store";

Vue.config.productionTip = false;
Vue.config.devtools = true;

import VueRouter from "vue-router";

Vue.component("vue-fontawesome", FontAwesomeIcon);
Vue.use(VueRouter);
Vue.use(Buefy, {
  defaultIconComponent: "vue-fontawesome",
  defaultIconPack: "fas",
});

Vue.config.productionTip = false;
Vue.config.devtools = true;

Wails.Init(() => {
  new Vue({
    render: (h) => h(App),
    router,
    store,
    mounted() {
      this.$router.replace("/");
    },
  }).$mount("#app");
});
