package store

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails"
)

var (
	ErrMustProvideDir = errors.New("must provide config dir")
	ErrMustContainDB  = errors.New("must contain .db")
	ErrValueEmpty     = errors.New("no data in entry")
	ErrFieldEmpty     = errors.New("no key provided")
	ErrQueryEmpty     = errors.New("no query provided")
)

type DB struct {
	db *storm.DB

	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

type Store interface {
	Set(interface{}) error
	Get(string, interface{}, interface{}) error
	Delete(interface{}) error
	Query(string, string, interface{}) error
}

// WailsInit performs setup when Wails is ready.
func (s *DB) WailsInit(runtime *wails.Runtime) error {
	s.Runtime = runtime
	s.Logger = s.Runtime.Log.New("Store")
	s.Logger.Info("Store initialized...")

	return nil
}

func Path(name string) string {
	config, err := os.UserConfigDir()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	return filepath.Join(config, "pod", name+".db")
}

func NewDB(name string) (*DB, error) {
	defaultDB := &DB{}

	path := Path(name)
	db, err := storm.Open(path)
	if err != nil {
		return defaultDB, err
	}

	return &DB{
		db: db,
	}, nil
}

func NewCustomDB(name string, folder string) (*DB, error) {
	defaultDB := &DB{}

	db, err := storm.Open(filepath.Join(folder, fmt.Sprintf("%s.db", name)))
	if err != nil {
		return defaultDB, err
	}

	return &DB{
		db: db,
	}, nil
}

func (s *DB) Close() error {
	return s.db.Close()
}

func (s *DB) Set(v interface{}) error {
	if v == nil {
		return ErrValueEmpty
	}

	return s.db.Save(v)
}

func (s *DB) Query(field, value string, to interface{}) error {
	if field == "" {
		return ErrFieldEmpty
	}

	if value == "" {
		return ErrValueEmpty
	}

	if to == nil {
		return ErrValueEmpty
	}

	regex := "^" + value

	return s.db.Select(q.Re(field, regex)).Find(to)
}

func (s *DB) List(to interface{}) error {
	return s.db.All(to)
}

func (s *DB) Get(fieldKey string, fieldValue interface{}, to interface{}) error {
	if fieldKey == "" {
		return ErrFieldEmpty
	}

	return s.db.One(fieldKey, fieldValue, to)
}

func (s *DB) Delete(v interface{}) error {
	return s.db.DeleteStruct(v)
}
