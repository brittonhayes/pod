package project

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	TableName = "projects"

	ErrFailedConnectDB = errors.New("failed to connect to database")
)

type Project struct {
	gorm.Model

	Name    string `json:"name"`
	Summary string `json:"summary"`
	Client  string `json:"client"`

	db string `gorm:"-"`
}

func NewProject(path string) *Project {
	return &Project{
		db: path,
	}
}

func (p *Project) With(name, summary, client string) *Project {
	p.Name = name
	p.Summary = summary
	p.Client = client
	return p
}

func (p *Project) setupLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{},
	)
}

func (p *Project) openDB() (*gorm.DB, error) {
	logger := p.setupLogger()
	db, err := gorm.Open(sqlite.Open(p.db), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger,
	})
	if err != nil {
		return nil, ErrFailedConnectDB
	}

	err = db.Model(p).AutoMigrate(p)
	return db, err
}

func (p *Project) Save() error {
	db, err := p.openDB()
	if err != nil {
		return err
	}

	db.Create(p)

	return nil
}

func (p *Project) Delete(name string) error {
	db, err := p.openDB()
	if err != nil {
		return err
	}

	tx := db.Where("name = ?", name).Delete(p)

	return tx.Error
}

func (p *Project) FindByName(name string, limit int) ([]Project, error) {
	db, err := p.openDB()
	if err != nil {
		return nil, err
	}

	var projects []Project
	tx := db.Model(p).Limit(limit).Where("name LIKE ?", name).Find(&projects)
	return projects, tx.Error
}

func (p *Project) List() ([]Project, error) {
	db, err := p.openDB()
	if err != nil {
		return nil, err
	}

	var projects []Project
	tx := db.Table(TableName).Select("*").Scan(&projects)
	return projects, tx.Error
}
