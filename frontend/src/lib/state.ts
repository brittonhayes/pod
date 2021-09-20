import { Project } from "./projects";
import runtime from "@wailsapp/runtime";

export class ProjectStore {
  store: any;
  state?: Project;

  constructor(name: string, defaultValue: Project) {
    this.store = runtime.Store.New(name, defaultValue);
  }

  set(project: Project) {
    runtime.Log.Info(`Setting project: ${project}`);
    this.store.set(project);
    this.state = project;
  }

  subscribe(project: Project) {
    this.store.subscribe(function(project: Project) {
      runtime.Log.Info(`Subscribing to project: ${project}`);
    });
  }

  update(project: Project): Project {
    runtime.Log.Info(`Updating project: ${project}`);
    this.state = project;
    return this.state;
  }
}
