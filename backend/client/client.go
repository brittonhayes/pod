package client

import (
	"os"
	"path/filepath"
	"time"

	"github.com/brittonhayes/pod/backend/store"
	"github.com/rs/zerolog/log"
)

type Client struct {
	ID          int               `storm:"id,increment" json:"id"`
	Name        string            `storm:"index,unique" json:"name"`
	Description string            `json:"description"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	Social      map[string]string `json:"social"`
	CreatedAt   time.Time         `json:"created_at"`

	db string `storm:"-"`
}

func Folder() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal().Err(err)
	}

	return filepath.Join(configDir, "pod", "clients")
}

func NewClient(path string) *Client {
	return &Client{
		db: path,
	}
}

func (c *Client) With(name, description, email, phone string) *Client {
	c.Name = name
	c.Description = description
	c.Email = email
	c.Phone = phone
	return c
}

func (c *Client) GetByID(id int, to Client) (bool, error) {
	db, err := store.NewDB(c.db)
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.Get("ID", id, &to)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) Save() (bool, error) {
	db, err := store.NewDB(c.db)
	if err != nil {
		return false, err
	}
	defer db.Close()

	c.CreatedAt = time.Now()
	err = db.Set(c)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) Delete() (bool, error) {
	db, err := store.NewDB(c.db)
	if err != nil {
		return false, err
	}
	defer db.Close()

	err = db.Delete(c)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) Query(field, value string, to []Client) (bool, error) {
	db, err := store.NewDB(c.db)
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
