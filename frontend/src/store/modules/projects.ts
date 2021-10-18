import { Project } from "@/types/project";
import { PropType } from "vue";
import {
  SET_ACTIVE,
  SET_LIST,
  SYNC_FORM,
  SAVE,
  REFRESH,
  SET_ENABLED,
  TOGGLE,
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
    [TOGGLE](state: any) {
      state.enabled = !state.enabled;
    },
    [SYNC_FORM](state: any, form: Project) {
      state.form = form;
    },
    [SAVE](state: any, form: Project) {
      window.backend.Storage.SaveProject(form)
        .then((res) => {
          state.active = res;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    [REFRESH](state: any) {
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
    [REFRESH](context: any) {
      context.commit(REFRESH);
    },
    [SAVE](context: any, form: Project) {
      context.commit(SAVE, form);
    },
  },
};
