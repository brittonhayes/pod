import Vue from "vue";
import Vuex from "vuex";

import { createLogger } from "vuex";

import { ProjectsModule } from "@/store/modules/projects";
import { ClientsModule } from "@/store/modules/clients";
import { ClipsModule } from "@/store/modules/clips";
import { RouterModule } from "@/store/modules/router";

Vue.use(Vuex);

export const store = new Vuex.Store({
  plugins: [createLogger()],
  modules: {
    projects: ProjectsModule,
    clients: ClientsModule,
    clips: ClipsModule,
    router: RouterModule,
  },
});
