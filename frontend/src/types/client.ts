import { Column } from "@/types/table";
import { GormModel } from "./gorm";

export interface Social {
  social_twitter?: string;
  social_instagram?: string;
  social_website?: string;
  social_facebook?: string;
  social_linkedin?: string;
}

export interface Client extends GormModel, Social {
  name: string;
  description?: string;
  email?: string;
  phone?: string;
}

export const ClientSort: Array<string> = ["id", "desc"];

export const ClientColumns: Array<Column> = [
  {
    label: "ID",
    field: "id",
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
    id: 1,
    name: "James",
    description: "Foo bar",
    email: "james@example.com",
    phone: "123-456-7890",
    created_at: new Date().toLocaleString(),
  },
  {
    id: 2,
    name: "Sarah",
    description: "Foo baz",
    email: "sarah@example.com",
    phone: "123-786-7890",
    created_at: new Date().toLocaleString(),
  },
];
