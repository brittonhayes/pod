package config

import (
	"os"
	"path/filepath"

	"github.com/wailsapp/wails"
)

type Config struct {
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

func NewConfig() *Config {
	return &Config{}
}

func Directory() string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, "pod")
}

func Setup() error {
	dbPath := Directory()
	err := os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
