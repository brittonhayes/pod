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
  faLink,
  faCut,
  faHeadphones,
  faFolder,
  faPencilAlt,
  faPlus,
  faRedo,
  faReply,
  faSearch,
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
  faCog,
  faLink,
  faCut,
  faHeadphones,
  faFolder,
  faPencilAlt,
  faPlus,
  faReply,
  faRedo,
  faSearch
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
import Clips from "./pages/Clips.vue";
import Settings from "./pages/Settings.vue";
// import the store
import store from "./store";

Vue.config.productionTip = false;
Vue.config.devtools = true;

const routes = [
  {
    path: "/",
    name: "pod",
    component: Home,
    meta: {
      title: "Pod 🌱",
      subtitle: "The audio professional's productivity center",
      icon: "leaf",
      iconSize: "is-large",
    },
  },
  {
    path: "/projects",
    name: "projects",
    component: Projects,
    meta: {
      title: "Projects 📝",
      subtitle: "Tasks, folders, and documents",
      icon: "tasks",
      iconSize: "is-medium",
    },
  },
  {
    path: "/clients",
    name: "clients",
    component: Clients,
    meta: {
      title: "Clients 👥",
      subtitle: "Customers and connections",
      icon: "address-book",
      iconSize: "is-medium",
    },
  },
  {
    path: "/clips",
    name: "clips",
    component: Clips,
    meta: {
      title: "Clips ✂️",
      subtitle: "Sound bytes, samples, and sessions",
      icon: "cut",
      iconSize: "is-medium",
    },
  },
  {
    path: "/settings",
    name: "settings",
    component: Settings,
    meta: {
      title: "Settings ⚙️",
      subtitle: "",
      icon: "cog",
      iconSize: "is-medium",
    },
  },
];

const router = new VueRouter({
  mode: "abstract",
  routes,
  linkExactActiveClass: "is-active", // active class for *exact* links.
});

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
