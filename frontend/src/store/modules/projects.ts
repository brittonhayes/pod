import { Project } from "@/types/project";
import { PropType } from "vue";
import {
  SET_ACTIVE,
  SET_LIST,
  SYNC_FORM,
  SUBMIT_FORM,
  UPDATE_FROM_DB,
  SET_ENABLED,
  TOGGLE_ENABLED,
  IS_ENABLED,
} from "@/store/mutations";

export const ProjectsModule = {
  namespaced: true,
  state: () => ({
    active: Object as PropType<Project>,
    projects: Array<Project>(),
    form: Object as PropType<Project>,
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
    [SYNC_FORM](state: any, form: Project) {
      state.form = form;
    },
    [SUBMIT_FORM](state: any, form: Project) {
      window.backend.Storage.SaveProject(form)
        .then((res) => {
          state.active = res;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    [UPDATE_FROM_DB](state: any) {
      window.backend.Storage.ListProjects()
        .then((res) => {
          state.projects = res;
        })
        .catch((err) => {
          console.error(err);
          state.projects = [{}];
        });
    },
    [SET_ACTIVE](state: any, id: number) {
      state.active = state.projects.find(
        (project: Project) => project.id === id
      );
    },
    [SET_LIST](state: any, payload: Array<Project>) {
      state.projects = payload;
    },
  },
  actions: {
    [UPDATE_FROM_DB](context: any) {
      context.commit(UPDATE_FROM_DB);
    },
    [SUBMIT_FORM](context: any, form: Project) {
      context.commit(SUBMIT_FORM, form);
    },
  },
};
