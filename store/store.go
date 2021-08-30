package store

import "fyne.io/fyne/v2"

type DB struct {
	driver fyne.App
}

type Store interface {
	Set(string, string)
	Get(string) string
}

func New(app fyne.App) *DB {
	return &DB{driver: app}
}

func (s *DB) Set(key, value string) {
	s.driver.Preferences().SetString(key, value)
}

func (s *DB) Get(key string) string {
	return s.driver.Preferences().String(key)
}
