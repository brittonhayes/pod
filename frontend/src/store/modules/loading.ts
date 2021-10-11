import { SET_ACTIVE } from "@/store/mutations";

export const LoadingModule = {
  namespaced: true,
  state: () => ({
    active: Boolean,
  }),
  mutations: {
    [SET_ACTIVE](state: any, active: Boolean) {
      state.active = active;
    },
  },
  actions: {
    [SET_ACTIVE](context: any, active: Boolean) {
      context.commit(SET_ACTIVE, active);
    },
  },
  getters: {},
};
