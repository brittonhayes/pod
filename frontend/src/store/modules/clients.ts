import { PropType } from "vue";

import { Client } from "@/types/client";
import {
  SET_ACTIVE,
  SET_LIST,
  UPDATE_FROM_DB,
} from "@/store/modules/mutations";

export const ClientsModule = {
  state: () => ({
    active: Object as PropType<Client>,
    clients: Array<Client>(),
  }),
  mutations: {
    [SET_ACTIVE](state: any, id: number) {
      state.active = state.clients.find((client: Client) => client.id === id);
    },
    [SET_LIST](state: any, payload: Array<Client>) {
      state.clients = payload;
    },
    [UPDATE_FROM_DB](state: any) {
      window.backend.Storage.ListClients()
        .then((res) => {
          state.clients = res;
        })
        .catch((err) => {
          console.error(err);
          state.clients = [{}];
        });
    },
  },
  actions: {},
  getters: {},
};
