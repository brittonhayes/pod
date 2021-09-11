package system

import (
	"os"
	"path/filepath"

	"github.com/brittonhayes/pod/project"
	"github.com/brittonhayes/pod/store"
)

func Initialize() error {
	projectPath := project.Folder()
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	clientPath := project.ClientFolder()
	err = os.MkdirAll(clientPath, os.ModePerm)
	if err != nil {
		return err
	}

	configPath := filepath.Dir(store.Path("pod"))
	err = os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
