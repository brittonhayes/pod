import { PropType } from "vue";
import { ToastProgrammatic as Toast } from "buefy";

import { Client } from "@/types/client";
import {
  SET_ACTIVE,
  SET_ENABLED,
  TOGGLE,
  SET_LIST,
  IS_ENABLED,
  REFRESH,
  SAVE,
  UPDATE,
  SEARCH,
} from "@/store/mutations";
import { Query } from "@/types/gorm";

export const ClientsModule = {
  namespaced: true,
  state: () => ({
    active: Object as PropType<Client>,
    clients: Array<Client>(),
    form: Object as PropType<Client>,
    enabled: false,
    limit: 25,
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
    [TOGGLE](state: any) {
      state.enabled = !state.enabled;
    },
    [SET_ACTIVE](state: any, id: number) {
      state.active = state.clients.find((client: Client) => client.id === id);
    },
    [SET_LIST](state: any, payload: Array<Client>) {
      state.clients = payload;
    },
    [SAVE](state: any, form: Client) {
      window.backend.Storage.SaveClient(form)
        .then((res) => {
          state.active = res;
        })
        .catch((err) => {
          console.error(err);
          Toast.open({
            message: "failed to submit",
            type: "is-danger",
            position: "is-bottom",
          });
        });
    },
    [UPDATE](state: any, form: Client) {
      window.backend.Storage.UpdateClient(form)
        .then((res) => {
          state.active = res;
          Toast.open({
            message: `Updated client`,
            type: "is-success",
            position: "is-bottom",
          });
        })
        .catch((err) => {
          console.error(err);
          Toast.open({
            message: "failed to submit",
            type: "is-warning",
            position: "is-bottom",
          });
        });
    },
    [SEARCH](state: any, query: Query) {
      window.backend.Storage.SearchClients(query.name, query.limit)
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
    [REFRESH](state: any) {
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
    [REFRESH](context: any) {
      context.commit(REFRESH);
    },
    [SAVE](context: any, form: Client) {
      context.commit(SAVE, form);
    },
    [SEARCH](context: any, query: Query) {
      context.commit(SAVE, query);
    },
    [UPDATE](context: any, form: Client) {
      context.commit(UPDATE, form);
    },
  },
};
