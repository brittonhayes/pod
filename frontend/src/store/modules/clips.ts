import { Mutator, CLIENTS, UPDATE_FROM_DB } from "@/store/mutations";

const mu = new Mutator(CLIENTS);

export const ClipsModule = {
  state: () => ({
    clips: Array<string>(),
  }),
  mutations: {
    [mu.Mutation(UPDATE_FROM_DB)](state: any) {
      window.backend.Storage.ListClips()
        .then((res) => {
          state.clips = res;
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
  actions: {
    [mu.Mutation(UPDATE_FROM_DB)](context: any) {
      context.commit(UPDATE_FROM_DB);
    },
  },
  getters: {},
};
