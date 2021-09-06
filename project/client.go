package project

import (
	"time"

	"github.com/brittonhayes/pod/input"
	"github.com/brittonhayes/pod/store"
	"github.com/google/uuid"
)

type Client struct {
	ID        uint32
	Name      string
	CreatedAt time.Time
}

func NewClient(name string) (*Client, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, ErrUUID
	}

	return &Client{
		ID:        id.ID(),
		Name:      name,
		CreatedAt: time.Now().Local(),
	}, nil
}

func (c *Client) Save() (bool, error) {
	db, err := store.NewDB(DBName)
	if err != nil {
		return false, ErrSave
	}

	key := input.StringToSnake(c.Name)
	err = db.Set(key, &c)
	if err != nil {
		return false, err
	}

	return true, nil
}
