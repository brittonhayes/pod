import { Project } from "@/types/project";
import { PropType } from "vue";
import {
  Mutator,
  SET_ACTIVE,
  SET_LIST,
  SYNC_FORM,
  SUBMIT_FORM,
  UPDATE_FROM_DB,
  SET_ENABLED,
  TOGGLE_ENABLED,
  IS_ENABLED,
  PROJECTS,
} from "@/store/mutations";

const mu = new Mutator(PROJECTS);

export const ProjectsModule = {
  state: () => ({
    active: Object as PropType<Project>,
    projects: Array<Project>(),
    form: Object as PropType<Project>,
    enabled: Boolean,
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
    [mu.Mutation(SYNC_FORM)](state: any, form: Project) {
      state.form = form;
    },
    [mu.Mutation(SUBMIT_FORM)](state: any, form: Project) {
      window.backend.Storage.SaveProject(form)
        .then((res) => {
          state.active = res;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    [mu.Mutation(UPDATE_FROM_DB)](state: any) {
      window.backend.Storage.ListProjects()
        .then((res) => {
          state.projects = res;
        })
        .catch((err) => {
          console.error(err);
          state.projects = [{}];
        });
    },
    [mu.Mutation(SET_ACTIVE)](state: any, id: number) {
      state.active = state.projects.find(
        (project: Project) => project.id === id
      );
    },
    [mu.Mutation(SET_LIST)](state: any, payload: Array<Project>) {
      state.projects = payload;
    },
  },
  actions: {
    [mu.Mutation(UPDATE_FROM_DB)](context: any) {
      context.commit(UPDATE_FROM_DB);
    },
    [mu.Mutation(SUBMIT_FORM)](context: any, form: Project) {
      context.commit(SUBMIT_FORM, form);
    },
  },
};
