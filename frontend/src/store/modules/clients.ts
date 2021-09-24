import { PropType } from "vue";

import { Client } from "@/types/client";
import { SET_ACTIVE, SET_LIST } from "@/store/modules/mutations";

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
  },
  actions: {},
  getters: {},
};
