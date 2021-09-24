package config

import (
	"os"
	"path/filepath"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/brittonhayes/pod/backend/store"
	"github.com/wailsapp/wails"
)

const (
	dbName = "pod"
)

type Config struct {
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

func NewConfig() *Config {
	return &Config{}
}

func Run() error {
	projectPath := project.Folder()
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	clientPath := client.Folder()
	err = os.MkdirAll(clientPath, os.ModePerm)
	if err != nil {
		return err
	}

	configDir, _ := os.UserConfigDir()

	dbPath := filepath.Dir(store.Path(configDir, dbName))
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
