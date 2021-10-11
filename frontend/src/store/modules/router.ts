import { SET_ACTIVE } from "@/store/mutations";

export const RouterModule = {
  namespaced: true,
  state: () => ({
    active: String,
  }),
  mutations: {
    [SET_ACTIVE](state: any, route: string) {
      state.active = route;
    },
  },
  actions: {
    [SET_ACTIVE](context: any, route: string) {
      context.commit(SET_ACTIVE, route);
    },
  },
  getters: {},
};
