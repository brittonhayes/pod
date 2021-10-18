export enum Namespace {
  Projects = "projects/",
  Clients = "clients/",
  Clips = "clips/",
  Router = "router/",
  Loading = "loading/",
}

type Mutation = string;
export const SET_ACTIVE: Mutation = "SET_ACTIVE";
export const SET_LIST: Mutation = "SET_LIST";
export const REFRESH: Mutation = "REFRESH";
export const SYNC_FORM: Mutation = "SYNC_FORM";
export const SAVE: Mutation = "SAVE";
export const UPDATE: Mutation = "UPDATE";
export const TOGGLE: Mutation = "TOGGLE";
export const SEARCH: Mutation = "SEARCH";
export const IS_ENABLED: Mutation = "IS_ENABLED";
export const SET_ENABLED: Mutation = "SET_ENABLED";
