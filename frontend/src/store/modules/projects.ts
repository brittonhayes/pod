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
} from "./mutations";

interface Form {
  Name: string;
  Summary: string;
  Client: string;
}

export const ProjectsModule = {
  state: () => ({
    active: Object as PropType<Project>,
    projects: Array<Project>(),
    form: Object as PropType<Form>,
    enabled: {
      type: Boolean,
      default: () => false,
    },
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
    [SYNC_FORM](state: any, form: Form) {
      state.form = form;
    },
    [SUBMIT_FORM](state: any, form: Form) {
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
    [SUBMIT_FORM](context: any, form: Form) {
      context.commit(SUBMIT_FORM, form);
    },
  },
};
