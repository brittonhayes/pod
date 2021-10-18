export interface GormModel {
  id: number;
  created_at?: string | null;
  deleted_at?: string | null;
  updated_at?: string | null;
}

export interface Query {
  name: string;
  limit: number;
}
