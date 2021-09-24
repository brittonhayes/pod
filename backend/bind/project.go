package bind

import (
	"fmt"

	"github.com/brittonhayes/pod/backend/project"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
)

type ProjectState struct {
	current *project.Project

	r      *wails.Runtime
	logger *wails.CustomLogger
	store  *wails.Store
}

// WailsInit is called when the component is being initialised
func (s *ProjectState) WailsInit(runtime *wails.Runtime) error {

	s.current = &project.Project{}
	s.r = runtime
	s.logger = runtime.Log.New("Project State")
	s.store = runtime.Store.New("project", &project.Project{})

	s.store.Subscribe(func(state *project.Project) {
		s.logger.Info(fmt.Sprintf("%#v", state))
	})

	return nil
}

func (s *ProjectState) Set(payload map[string]interface{}) {

	p := &project.Project{}
	err := mapstructure.Decode(p, &payload)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	s.logger.Info(fmt.Sprintf("Setting value: %s", p.Name))
	s.current = p
	s.store.Set(p)
}

func (s *ProjectState) Value() *project.Project {
	s.store.Update(func(current *project.Project) *project.Project {
		return s.current
	})

	return s.current
}
