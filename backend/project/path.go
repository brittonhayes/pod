package project

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func Folder() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal().Err(err)
	}

	return filepath.Join(configDir, "pod", "projects")
}
