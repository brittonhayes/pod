import { UPDATE_FROM_DB } from "@/store/mutations";

export const ClipsModule = {
  namespaced: true,
  state: () => ({
    clips: Array<string>(),
  }),
  mutations: {
    [UPDATE_FROM_DB](state: any) {
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
    [UPDATE_FROM_DB](context: any) {
      context.commit(UPDATE_FROM_DB);
    },
  },
  getters: {},
};
