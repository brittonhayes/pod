import VueRouter from "vue-router";

import Home from "@/pages/Home.vue";
import Clients from "@/pages/Clients.vue";
import Profile from "@/pages/Profile.vue";
import Projects from "@/pages/Projects.vue";
import Clips from "@/pages/Clips.vue";
import Settings from "@/pages/Settings.vue";

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
      title: "Clients 📇",
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
      subtitle: "Customize and configure pod",
      icon: "cog",
      iconSize: "is-medium",
    },
  },
  {
    path: "/profile/:id",
    name: "profile",
    component: Profile,
    meta: {
      title: "Profile 📇",
      subtitle: "Customers and connections",
      icon: "address-book",
      iconSize: "is-medium",
    },
  },
];

const router = new VueRouter({
  mode: "abstract",
  routes,
  linkExactActiveClass: "is-active", // active class for *exact* links.
});

export default router;
