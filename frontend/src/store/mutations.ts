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
export const UPDATE_FROM_DB: Mutation = "UPDATE_FROM_DB";
export const SYNC_FORM: Mutation = "SYNC_FORM";
export const SUBMIT_FORM: Mutation = "SUBMIT_FORM";
export const UPDATE_ITEM: Mutation = "UPDATE_ITEM";
export const TOGGLE_ENABLED: Mutation = "TOGGLE_ENABLED";
export const IS_ENABLED: Mutation = "IS_ENABLED";
export const SET_ENABLED: Mutation = "SET_ENABLED";
