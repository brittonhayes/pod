import { PropType } from "vue";

import { Client } from "@/types/client";
import {
  SET_ACTIVE,
  SET_ENABLED,
  TOGGLE_ENABLED,
  SET_LIST,
  IS_ENABLED,
  UPDATE_FROM_DB,
  SUBMIT_FORM,
} from "@/store/mutations";

export const ClientsModule = {
  namespaced: true,
  state: () => ({
    active: Object as PropType<Client>,
    clients: Array<Client>(),
    form: Object as PropType<Client>,
    enabled: false,
  }),
  getters: {
    [IS_ENABLED](state: any): Boolean {
      return state.enabled;
    },
  },
  mutations: {
    [SET_ENABLED](state: any, value: Boolean) {
      state.enabled = value;
    },
    [TOGGLE_ENABLED](state: any) {
      state.enabled = !state.enabled;
    },
    [SET_ACTIVE](state: any, id: number) {
      state.active = state.clients.find((client: Client) => client.ID === id);
    },
    [SET_LIST](state: any, payload: Array<Client>) {
      state.clients = payload;
    },
    [SUBMIT_FORM](state: any, form: Client) {
      window.backend.Storage.SaveClient(form)
        .then((res) => {
          state.active = res;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    [UPDATE_FROM_DB](state: any) {
      window.backend.Storage.ListClients()
        .then((res) => {
          if (res !== null) {
            state.clients = res;
          }
        })
        .catch((err) => {
          console.error(err);
          state.clients = [{}];
        });
    },
  },
  actions: {
    [UPDATE_FROM_DB](context: any) {
      context.commit(UPDATE_FROM_DB);
    },
    [SUBMIT_FORM](context: any, form: Client) {
      context.commit(SUBMIT_FORM, form);
    },
  },
};
