import Vue from "vue";
import Vuex from "vuex";
import { Project } from "@/types/project";
import { Client } from "@/types/client";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    route: {
      type: String,
    },
    projects: Array<Project>(),
    clients: Array<Client>(),
    clips: Array<Object>(),
  },
  getters: {
    getActiveRoute(state) {
      return state.route;
    },
  },
  actions: {},
  mutations: {
    setActiveRoute(state, route) {
      state.route = route;
    },
  },
});
