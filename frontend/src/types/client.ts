import { Column } from "@/types/table";

export interface Client {
  id: number;
  name: string;
  description?: string;
  email?: string;
  phone?: string;
  social?: Object;
  created_at?: string;
}

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

export const ClientSort: Array<string> = ["created_at", "asc"];

export const ClientColumns: Array<Column> = [
  {
    field: "id",
    label: "id",
    numeric: true,
    sortable: true,
    visible: false,
  },
  {
    field: "name",
    label: "Client",
    sortable: true,
    searchable: true,
  },
  {
    field: "email",
    label: "email",
    sortable: true,
  },
  {
    field: "phone",
    label: "Phone",
    sortable: true,
  },
  {
    field: "created_at",
    label: "Created",
    sortable: true,
  },
];
