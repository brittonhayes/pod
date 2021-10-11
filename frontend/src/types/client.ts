import { Column } from "@/types/table";
import { GormModel } from "./gorm";

export interface Client extends GormModel {
  name: string;
  description?: string;
  email?: string;
  phone?: string;
  social?: Object;
}

export const ClientSort: Array<string> = ["id", "desc"];

export const ClientColumns: Array<Column> = [
  {
    label: "ID",
    field: "ID",
    width: 10,
    numeric: true,
    sortable: true,
  },
  {
    field: "name",
    label: "Name",
    sortable: true,
    searchable: false,
  },
  {
    field: "email",
    label: "Email",
    sortable: true,
  },
  {
    field: "phone",
    label: "Phone",
    sortable: true,
  },
  {
    field: "CreatedAt",
    label: "Created",
    sortable: true,
    visible: false,
  },
];

export const DefaultClients: Array<Client> = [
  {
    ID: 1,
    name: "James",
    description: "Foo bar",
    email: "james@example.com",
    phone: "123-456-7890",
    CreatedAt: new Date().toLocaleString(),
  },
  {
    ID: 2,
    name: "Sarah",
    description: "Foo baz",
    email: "sarah@example.com",
    phone: "123-786-7890",
    CreatedAt: new Date().toLocaleString(),
  },
];
