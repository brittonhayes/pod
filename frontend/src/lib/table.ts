export type Column = {
  field: string;
  label?: string;
  "custom-key"?: string;
  meta?: any;
  width?: number | string;
  numeric?: boolean;
  centered?: boolean;
  sortable?: boolean;
  visible?: boolean;
  searchable?: boolean;
  subheading?: string | number;
  sticky?: boolean;
  "header-selectable"?: boolean;
  "header-class"?: string;
  "cell-class"?: string;
};
