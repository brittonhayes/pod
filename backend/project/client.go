package project

import (
	"errors"
	"fmt"
	"time"

	"github.com/brittonhayes/pod/backend/store"
)

type Contact struct {
	Email  string
	Phone  string
	Social map[string]string
}
type Client struct {
	ID          int    `storm:"id,increment"`
	Name        string `storm:"index,unique"`
	Description string
	Contact     Contact
	CreatedAt   time.Time
}

func NewClient(name string) *Client {
	return &Client{
		Name:        name,
		Description: "",
	}
}

func (c *Client) Save() (bool, error) {
	db, err := store.NewDB(DBName)
	if err != nil {
		return false, ErrSave
	}
	defer db.Close()

	c.CreatedAt = time.Now()
	err = db.Set(c)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Client) Delete() (bool, error) {
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

func (c *Client) Query(field, value string, to []Client) (bool, error) {
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