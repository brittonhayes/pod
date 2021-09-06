package store

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/asdine/storm/v3"
)

var (
	ErrMustProvideDir = errors.New("must provide config dir")
	ErrMustContainDB  = errors.New("must contain .db")
	ErrValueEmpty     = errors.New("no data in entry")
	ErrKeyEmpty       = errors.New("no key provided")
)

type DB struct {
	bucket string
	db     *storm.DB
}

type Store interface {
	Set(string, interface{}) error
	Get(string, interface{}) error
	Delete(string) error
}

func (db *DB) Path() (string, error) {
	if db.bucket == "" {
		return ErrMustProvideDir.Error(), ErrMustProvideDir
	}

	config, err := os.UserConfigDir()
	if err != nil {
		return ErrMustContainDB.Error(), err
	}

	return filepath.Join(config, db.bucket, db.bucket+".db"), nil
}

func NewDB(name string) (*DB, error) {

	defaultDB := &DB{}

	config, err := os.UserConfigDir()
	if err != nil {
		return defaultDB, err
	}

	dbPath := filepath.Join(config, name, name+".db")
	db, err := storm.Open(dbPath)
	if err != nil {
		return defaultDB, err
	}

	return &DB{
		bucket: name,
		db:     db,
	}, nil
}

func (s *DB) Set(k string, v interface{}) error {
	if k == "" {
		return ErrKeyEmpty
	}

	if v == nil {
		return ErrValueEmpty
	}

	return s.db.Set(s.bucket, k, &v)
}

func (s *DB) Get(k string, v interface{}) error {
	if k == "" {
		return ErrKeyEmpty
	}

	return s.db.Get(s.bucket, k, &v)
}

func (s *DB) Delete(k string) error {
	if k == "" {
		return ErrKeyEmpty
	}

	return s.db.Delete(s.bucket, k)
}
