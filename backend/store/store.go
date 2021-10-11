package store

import (
	"errors"
	"os"
	"strings"

	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

var (
	ErrMustProvideDir  = errors.New("must provide config dir")
	ErrMustContainDB   = errors.New("must contain .db")
	ErrValueEmpty      = errors.New("no data in entry")
	ErrFieldEmpty      = errors.New("no key provided")
	ErrQueryEmpty      = errors.New("no query provided")
	ErrFailedConnectDB = errors.New("failed to connect to database")
)

func setupLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func New(path string, migrations ...interface{}) (*gorm.DB, error) {
	logger := setupLogger()
	if !strings.HasSuffix(path, ".db") {
		path += ".db"
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger,
	})
	if err != nil {
		return nil, ErrFailedConnectDB
	}

	err = db.AutoMigrate(migrations...)

	return db, err
}
