interface Storage {
	FindClient(...args : any[]):Promise<any>
	FindProject(...args : any[]):Promise<any>
	ListClients(...args : any[]):Promise<any>
	ListClips(...args : any[]):Promise<any>
	ListProjects(...args : any[]):Promise<any>
	SaveClient(...args : any[]):Promise<any>
	SaveProject(...args : any[]):Promise<any>
	WailsInit(...args : any[]):Promise<any>
}
interface ProjectState {
	Set(...args : any[]):Promise<any>
	Value(...args : any[]):Promise<any>
	WailsInit(...args : any[]):Promise<any>
}
interface Config {
}

interface Backend {
	Config: Config
	Storage: Storage
	ProjectState: ProjectState
}

declare global {
	interface Window {
		backend: Backend;
	}
}
export {};