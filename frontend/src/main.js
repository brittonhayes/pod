import "core-js/stable";
import "regenerator-runtime/runtime";
import * as Wails from "@wailsapp/runtime";
import Vue from "vue";

import "buefy/dist/buefy.css";

import { library } from "@fortawesome/fontawesome-svg-core";
// internal icons
import {
  faCheck,
  faCheckCircle,
  faInfoCircle,
  faExclamationTriangle,
  faExclamationCircle,
  faArrowUp,
  faAngleRight,
  faAngleLeft,
  faAngleDown,
  faEye,
  faEyeSlash,
  faCaretDown,
  faCaretUp,
  faUpload,
  faUser,
  faHome,
  faTree,
  faTasks,
  faProjectDiagram,
  faLeaf,
  faAddressBook,
  faVolumeUp,
  faCog,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

library.add(
  faCheck,
  faCheckCircle,
  faInfoCircle,
  faExclamationTriangle,
  faExclamationCircle,
  faArrowUp,
  faAngleRight,
  faAngleLeft,
  faAngleDown,
  faEye,
  faEyeSlash,
  faCaretDown,
  faCaretUp,
  faUpload,
  faHome,
  faUser,
  faTree,
  faTasks,
  faLeaf,
  faAddressBook,
  faProjectDiagram,
  faVolumeUp,
  faCog
);

Vue.component("vue-fontawesome", FontAwesomeIcon);
Vue.config.productionTip = false;
Vue.config.devtools = true;

import VueRouter from "vue-router";
Vue.use(VueRouter);

import Buefy from "buefy";
Vue.use(Buefy, {
  defaultIconComponent: "vue-fontawesome",
  defaultIconPack: "fas",
});

import App from "./App.vue";
import Home from "./pages/Home.vue";
import Clients from "./pages/Clients.vue";
import Projects from "./pages/Projects.vue";
import Settings from "./pages/Settings.vue";

const routes = [
  { path: "/", name: "home", component: Home },
  { path: "/clients", name: "clients", component: Clients },
  { path: "/projects", name: "projects", component: Projects },
  { path: "/settings", name: "settings", component: Settings },
];

const router = new VueRouter({
  mode: "history",
  routes,
});

Wails.Init(() => {
  new Vue({
    router,
    render: (h) => h(App),
  }).$mount("#app");
});
