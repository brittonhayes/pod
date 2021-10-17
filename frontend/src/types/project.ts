import { Column } from "@/types/table";
import { Client } from "./client";
import { GormModel } from "./gorm";

export interface Project extends GormModel {
  name: string;
  summary?: string;

  ClientID?: number;
  client?: Client;
}

export const ProjectSort: Array<string> = ["created_at", "asc"];

export const ProjectColumns: Array<Column> = [
  {
    field: "id",
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
    field: "created_at",
    label: "Created",
    visible: false,
    sortable: true,
    width: 200,
  },
];

export const DefaultProjects: Array<Project> = [
  {
    id: 1,
    name: "First Project",
    summary: "Foo bar",
    client: { id: 1, name: "joey" },
    created_at: new Date().toLocaleString(),
  },
  {
    id: 2,
    name: "Second Project",
    summary: "Foo baz",
    client: { id: 2, name: "sarah" },
    created_at: new Date().toLocaleString(),
  },
];
