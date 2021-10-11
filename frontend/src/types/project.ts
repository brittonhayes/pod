import { Column } from "@/types/table";
import { Client } from "./client";
import { GormModel } from "./gorm";

export interface Project extends GormModel {
  name: string;
  summary?: string;

  ClientID?: number;
  client?: Client;
}

export const ProjectSort: Array<string> = ["CreatedAt", "asc"];

export const ProjectColumns: Array<Column> = [
  {
    field: "ID",
    label: "ID",
    width: 10,
    numeric: true,
    sortable: true,
  },
  {
    field: "Name",
    label: "Project",
    sortable: true,
  },
  {
    field: "CreatedAt",
    label: "Created",
    visible: false,
    sortable: true,
    width: 200,
  },
];

export const DefaultProjects: Array<Project> = [
  {
    ID: 1,
    name: "First Project",
    summary: "Foo bar",
    client: { ID: 1, name: "joey" },
    CreatedAt: new Date().toLocaleString(),
  },
  {
    ID: 2,
    name: "Second Project",
    summary: "Foo baz",
    client: { ID: 2, name: "sarah" },
    CreatedAt: new Date().toLocaleString(),
  },
];
