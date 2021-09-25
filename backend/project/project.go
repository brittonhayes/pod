package project

import (
	"time"

	"github.com/brittonhayes/pod/backend/store"
)

type Project struct {
	ID        uint32 `storm:"id,increment" json:"id"`
	Name      string `storm:"index,unique" json:"name"`
	Summary   string `json:"summary"`
	Client    string `json:"client"`
	CreatedAt string `json:"created_at"`

	db string `storm:"-"`
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

func (p *Project) Save() (bool, error) {
	db, err := store.NewDB(p.db)
	if err != nil {
		return false, err
	}
	defer db.Close()

	p.CreatedAt = time.Now().Format(time.RFC1123)
	err = db.Set(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Project) Delete() (bool, error) {
	db, err := store.NewDB(p.db)
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
	db, err := store.NewDB(p.db)
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
	db, err := store.NewDB(p.db)
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
