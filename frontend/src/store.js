import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    config: {
      appName: "pod",
    },
    projects: [],
    clients: [],
    settings: {},
    clips: [],
  },
  getters: {
    getActiveRoute(state) {
      return state.route;
    },
    config(state) {
      return state.config;
    },
  },
  actions: {},
  mutations: {
    setActiveRoute(state, route) {
      state.route = route;
    },
  },
});

export default store;
