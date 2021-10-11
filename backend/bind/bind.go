package bind

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/config"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/wailsapp/wails"
)

var (
	ErrPayloadNil = errors.New("submitted payload was nil")
)

type Storage struct {
	r      *wails.Runtime
	logger *wails.CustomLogger

	dbPath string
}

func (s *Storage) WailsInit(runtime *wails.Runtime) error {
	s.r = runtime
	s.logger = s.r.Log.New("Storage")
	s.dbPath = filepath.Join(config.Directory(), "pod.db")

	config.Setup()
	return nil
}

func (s *Storage) SaveClient(payload map[string]interface{}) error {
	if payload == nil {
		return ErrPayloadNil
	}

	c := client.New(s.dbPath)
	return c.SaveJSON(payload)
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
	c, err := client.New(s.dbPath).List()
	return c, err
}

func (s *Storage) ListProjects() ([]project.Project, error) {
	p, err := project.New(s.dbPath).List()
	return p, err
}

func (s *Storage) FindProject(name string, limit int) ([]project.Project, error) {
	p, err := project.New(s.dbPath).FindByName(name, limit)
	return p, err
}

func (s *Storage) FindClient(id int) (*client.Client, error) {
	return client.New(s.dbPath).GetByID(id)
}

func (s *Storage) SaveProject(payload map[string]interface{}) (*project.Project, error) {
	if payload == nil {
		return nil, ErrPayloadNil
	}

	p := project.New(s.dbPath)

	return p, p.SaveJSON(payload)
}
