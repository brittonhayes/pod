interface Config {
  WailsInit(...args: any[]): Promise<any>;
}
interface Storage {
  ListProjects(...args: any[]): Promise<any>;
  QueryClients(...args: any[]): Promise<any>;
  QueryProjects(...args: any[]): Promise<any>;
  SaveClient(...args: any[]): Promise<any>;
  SaveProject(...args: any[]): Promise<any>;
}

interface Backend {
  Config: Config;
  Storage: Storage;
}

declare global {
  interface Window {
    backend: Backend;
  }
}
export {};
