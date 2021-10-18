import { REFRESH } from "@/store/mutations";

export const ClipsModule = {
  namespaced: true,
  state: () => ({
    clips: Array<string>(),
  }),
  mutations: {
    [REFRESH](state: any) {
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
    [REFRESH](context: any) {
      context.commit(REFRESH);
    },
  },
  getters: {},
};
