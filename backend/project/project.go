package project

import (
	"errors"
	"time"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/store"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

var (
	TableName = "projects"

	ErrMissingDB = errors.New("no database path provided for project")
)

type Project struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Name     string         `json:"name"`
	Summary  string         `json:"summary"`
	ClientID uint           `json:"client_id"`
	Client   *client.Client `json:"client"`

	db string `gorm:"-"`
}

func New(path string) *Project {
	return &Project{
		db:     path,
		Client: &client.Client{},
	}
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	if p.db == "" {
		return ErrMissingDB
	}
	return
}

func (p *Project) With(name string, summary string, client *client.Client) *Project {
	p.Name = name
	p.Summary = summary
	p.Client = client
	return p
}

func (p *Project) Save() error {
	db, err := store.New(p.db, &Project{})
	if err != nil {
		return err
	}

	tx := db.Model(&p).Create(p)
	return tx.Error
}

func (p *Project) SaveJSON(payload map[string]interface{}) error {
	db, err := store.New(p.db, &Project{})
	if err != nil {
		return err
	}

	err = mapstructure.Decode(payload, &p)
	if err != nil {
		return err
	}

	tx := db.Model(&p).Save(p)
	return tx.Error
}

func (p *Project) Delete(name string) error {
	db, err := store.New(p.db, &Project{})
	if err != nil {
		return err
	}

	tx := db.Where("name = ?", name).Delete(p)
	return tx.Error
}

func (p *Project) FindByName(name string, limit int) ([]Project, error) {
	db, err := store.New(p.db, &Project{})
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
