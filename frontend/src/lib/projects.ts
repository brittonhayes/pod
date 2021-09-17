import { Column } from "./table";

export interface Project {
  id: number;
  name: string;
  summary?: string;
  client?: string;
  created_at?: string;
}

export const DefaultProjects: Array<Project> = [
  {
    id: 1,
    name: "First Project",
    summary: "Foo bar",
    client: "James",
    created_at: new Date().toLocaleString(),
  },
  {
    id: 2,
    name: "Second Project",
    summary: "Foo baz",
    client: "Sarah",
    created_at: new Date().toLocaleString(),
  },
];

export const ProjectSort: Array<string> = ["created_at", "desc"];

export const ProjectColumns: Array<Column> = [
  {
    field: "id",
    label: "ID",
    numeric: true,
  },
  {
    field: "name",
    label: "Project",
    searchable: true,
    sortable: true,
  },
  {
    field: "client",
    label: "Client",
    searchable: true,
    sortable: true,
  },
  {
    field: "created_at",
    label: "Created",
    sortable: true,
    width: 200,
  },
];
