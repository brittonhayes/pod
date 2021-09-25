import { PropType } from "vue";

import { Client } from "@/types/client";
import {
  CLIENTS,
  SET_ACTIVE,
  SET_ENABLED,
  TOGGLE_ENABLED,
  SET_LIST,
  IS_ENABLED,
  UPDATE_FROM_DB,
  SUBMIT_FORM,
} from "@/store/mutations";

import { Mutator } from "@/store/mutations";

const mu = new Mutator(CLIENTS);

export const ClientsModule = {
  state: () => ({
    active: Object as PropType<Client>,
    clients: Array<Client>(),
    form: Object as PropType<Client>,
    enabled: false,
  }),
  getters: {
    [mu.Mutation(IS_ENABLED)](state: any): Boolean {
      return state.enabled;
    },
  },
  mutations: {
    [mu.Mutation(SET_ENABLED)](state: any, value: Boolean) {
      state.enabled = value;
    },
    [mu.Mutation(TOGGLE_ENABLED)](state: any) {
      state.enabled = !state.enabled;
    },
    [mu.Mutation(SET_ACTIVE)](state: any, id: number) {
      state.active = state.clients.find((client: Client) => client.id === id);
    },
    [mu.Mutation(SET_LIST)](state: any, payload: Array<Client>) {
      state.clients = payload;
    },
    [mu.Mutation(SUBMIT_FORM)](state: any, form: Client) {
      window.backend.Storage.SaveClient(form)
        .then((res) => {
          state.active = res;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    [mu.Mutation(UPDATE_FROM_DB)](state: any) {
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
    [mu.Mutation(UPDATE_FROM_DB)](context: any) {
      context.commit(mu.Mutation(UPDATE_FROM_DB));
    },
    [mu.Mutation(SUBMIT_FORM)](context: any, form: Client) {
      context.commit(SUBMIT_FORM, form);
    },
  },
};
