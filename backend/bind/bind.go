package bind

import (
	"strings"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/brittonhayes/pod/backend/store"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
)

type Storage struct {
	r      *wails.Runtime
	logger *wails.CustomLogger

	dbPath string
}

func (s *Storage) WailsInit(runtime *wails.Runtime) error {
	s.r = runtime
	s.logger = s.r.Log.New("Storage")

	s.dbPath = store.Path("pod", "pod.db")
	return nil
}

func (s *Storage) SaveClient(payload map[string]interface{}) (*client.Client, error) {
	c := client.NewClient(s.dbPath)
	err := mapstructure.Decode(payload, &c)
	if err != nil {
		return c, err
	}

	_, err = c.Save()
	if err != nil {
		s.logger.Error(err.Error())
		return c, err
	}

	return c, nil
}

func (s *Storage) QueryClients(field, value string) ([]client.Client, error) {
	c := client.NewClient(s.dbPath)
	results := []client.Client{}
	_, err := c.Query(strings.ToTitle(field), value, results)
	if err != nil {
		s.logger.Error(err.Error())
		return []client.Client{}, err
	}

	return results, nil
}

func (s *Storage) ListProjects() ([]project.Project, error) {

	results := []project.Project{}
	db, err := store.NewDB(s.dbPath)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	err = db.List(&results)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return results, err
}

func (s *Storage) QueryProjects(field, value string) ([]project.Project, error) {

	p := project.NewProject(s.dbPath)
	results := []project.Project{}
	_, err := p.Query(strings.ToTitle(field), value, results)
	if err != nil {
		s.logger.Error(err.Error())
		return []project.Project{}, err
	}

	return results, nil
}

func (s *Storage) SaveProject(payload map[string]interface{}) (*project.Project, error) {
	p := project.NewProject(s.dbPath)
	err := mapstructure.Decode(payload, &p)
	if err != nil {
		s.logger.Error(err.Error())
		return p, err
	}

	_, err = p.Save()
	if err != nil {
		s.logger.Error(err.Error())
		return p, err
	}

	return p, nil
}
