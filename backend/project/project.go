package project

import (
	"errors"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/store"
	"gorm.io/gorm"
)

var (
	TableName = "projects"

	ErrMissingDB = errors.New("no database path provided for project")
)

type Project struct {
	gorm.Model

	Name    string
	Summary string

	ClientID int
	Client   client.Client

	db string `gorm:"-"`
}

func New(path string) *Project {
	return &Project{
		db: path,
	}
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	if p.db == "" {
		return ErrMissingDB
	}
	return
}

func (p *Project) With(name string, summary string, client client.Client) *Project {
	p.Name = name
	p.Summary = summary
	p.Client = client
	return p
}

func (p *Project) Save() error {
	db, err := store.New(p.db, p)
	if err != nil {
		return err
	}

	tx := db.Model(&p).Create(p)
	return tx.Error
}

func (p *Project) SaveJSON(payload map[string]interface{}) error {
	db, err := store.New(p.db, p)
	if err != nil {
		return err
	}

	tx := db.Model(&p).Create(payload)
	return tx.Error
}

func (p *Project) Delete(name string) error {
	db, err := store.New(p.db, p)
	if err != nil {
		return err
	}

	tx := db.Where("name = ?", name).Delete(p)
	return tx.Error
}

func (p *Project) FindByName(name string, limit int) ([]Project, error) {
	db, err := store.New(p.db, p)
	if err != nil {
		return nil, err
	}

	var projects []Project
	tx := db.Model(p).Limit(limit).Where("name LIKE ?", name).Find(&projects)
	return projects, tx.Error
}

func (p *Project) List() ([]Project, error) {
	db, err := store.New(p.db, &Project{})
	if err != nil {
		return nil, err
	}

	var projects []Project
	tx := db.Table(TableName).Select("*").Scan(&projects)
	return projects, tx.Error
}
