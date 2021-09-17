export interface Client {
  ID: number;
  Name: string;
  Description?: string;
  Email?: string;
  Phone?: string;
  Social?: Object;
}

export const DefaultClients: Array<Client> = [
  {
    ID: 1,
    Name: "James",
    Description: "Foo bar",
  },
  {
    ID: 2,
    Name: "Sarah",
    Description: "Foo baz",
    Email: "sarah@example.com",
  },
];
