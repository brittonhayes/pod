package project

import (
	"time"

	"github.com/brittonhayes/pod/backend/store"
)

const (
	dbName = "project"
)

type Project struct {
	ID        uint32    `storm:"id,increment" json:"id"`
	Name      string    `storm:"index,unique" json:"name"`
	Summary   string    `json:"summary"`
	Client    string    `json:"client"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) With(name, summary, client string) *Project {
	p.Name = name
	p.Summary = summary
	p.Client = client
	return p
}

func (p *Project) Save() (bool, error) {
	db, err := store.NewDB(dbName)
	if err != nil {
		return false, err
	}
	defer db.Close()

	p.CreatedAt = time.Now()
	err = db.Set(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Project) Delete() (bool, error) {
	db, err := store.NewDB(dbName)
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.Delete(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Project) Query(field, value string, to []Project) (bool, error) {
	db, err := store.NewDB(dbName)
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.Query(field, value, &to)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Project) List(to []Project) (bool, error) {
	db, err := store.NewDB(dbName)
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.List(&to)
	if err != nil {
		return false, err
	}

	return true, nil
}
