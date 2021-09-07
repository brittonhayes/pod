package project

import (
	"errors"
	"fmt"
	"time"

	"github.com/brittonhayes/pod/input"
	"github.com/brittonhayes/pod/store"
	"github.com/google/uuid"
)

const (
	DBName = "pod"
)

type Project struct {
	ID        uint32
	Name      string
	Client    *Client
	CreatedAt time.Time
}

func NewProject(name string) (*Project, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, ErrUUID
	}

	return &Project{
		ID:        id.ID(),
		Name:      input.StringToSnake(name),
		Client:    &Client{},
		CreatedAt: time.Now().Local(),
	}, nil
}

func (p *Project) Save() (bool, error) {
	db, err := store.NewDB(DBName)
	if err != nil {
		e := fmt.Sprintf("%s - %s", err.Error(), ErrSave.Error())
		return false, errors.New(e)
	}
	defer db.Close()

	key := input.StringToSnake(p.Name)
	err = db.Set(key, &p)
	if err != nil {
		return false, err
	}

	return true, nil
}
