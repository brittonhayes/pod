package bind

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
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

	s.dbPath = store.Path("pod", "pod")
	return nil
}

func (s *Storage) SaveClient(payload map[string]interface{}) (*client.Client, error) {
	if payload == nil {
		return nil, errors.New("submitted client was nil")
	}

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

func (s *Storage) ListClips() ([]string, error) {
	home, _ := os.UserHomeDir()
	clipsDir := filepath.Join(home, "pod", "clips")

	err := os.MkdirAll(clipsDir, os.ModePerm)
	if err != nil {
		return []string{""}, err
	}

	matches, err := fs.Glob(os.DirFS(clipsDir), "*.wav")
	if err != nil {
		return []string{""}, err
	}

	if len(matches) == 0 {
		return []string{""}, nil
	}

	return matches, nil
}

func (s *Storage) ListClients() ([]client.Client, error) {

	results := []client.Client{}
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

// func (s *Storage) QueryProjects(field, value string) ([]project.Project, error) {

// 	p := project.NewProject(s.dbPath)
// 	results := []project.Project{}
// 	_, err := p.FindByName(strings.ToTitle(field), value)
// 	if err != nil {
// 		s.logger.Error(err.Error())
// 		return []project.Project{}, err
// 	}

// 	return results, nil
// }

func (s *Storage) SaveProject(payload map[string]interface{}) (*project.Project, error) {
	if payload == nil {
		err := errors.New("submitted project was nil")
		s.logger.Error(err.Error())
		return nil, err
	}

	p := project.NewProject(s.dbPath)
	err := mapstructure.Decode(payload, &p)
	if err != nil {
		s.logger.Error(err.Error())
		return p, err
	}

	err = p.Save()
	if err != nil {
		s.logger.Error(err.Error())
		return p, err
	}

	return p, nil
}
