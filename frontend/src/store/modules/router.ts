import { Mutator, ROUTER, SET_ACTIVE } from "@/store/mutations";

const mu = new Mutator(ROUTER);

export const RouterModule = {
  state: () => ({
    active: String,
  }),
  mutations: {
    [mu.Mutation(SET_ACTIVE)](state: any, route: string) {
      state.active = route;
    },
  },
  actions: {
    [mu.Mutation(SET_ACTIVE)](context: any, route: string) {
      context.commit(mu.Mutation(SET_ACTIVE), route);
    },
  },
  getters: {},
};
