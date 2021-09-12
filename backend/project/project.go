package project

import (
	"errors"
	"fmt"
	"time"

	"github.com/brittonhayes/pod/backend/store"
)

const (
	DBName = "pod"
)

type Project struct {
	ID        uint32 `storm:"id,increment"`
	Name      string `storm:"index,unique"`
	Summary   string
	Client    *Client
	CreatedAt time.Time
}

func NewProject(name string) *Project {
	return &Project{
		Name:    name,
		Client:  &Client{},
		Summary: "",
	}
}

func (p *Project) Save() (bool, error) {
	db, err := store.NewDB(DBName)
	if err != nil {
		e := fmt.Sprintf("%s - %s", err.Error(), ErrSave.Error())
		return false, errors.New(e)
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
	db, err := store.NewDB(DBName)
	if err != nil {
		e := fmt.Sprintf("%s - %s", err.Error(), ErrSave.Error())
		return false, errors.New(e)
	}
	defer db.Close()

	err = db.Delete(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Project) Query(field, value string, to []Project) (bool, error) {
	db, err := store.NewDB(DBName)
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
