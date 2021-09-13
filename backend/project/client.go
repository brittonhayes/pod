package project

import (
	"errors"
	"fmt"
	"time"

	"github.com/brittonhayes/pod/backend/store"
)

type Client struct {
	ID          int               `storm:"id,increment" json:"id"`
	Name        string            `storm:"index,unique" json:"name"`
	Description string            `json:"description"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	Social      map[string]string `json:"social"`
	CreatedAt   time.Time         `json:"created_at"`
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) With(name, description, email, phone string) *Client {
	c.Name = name
	c.Description = description
	c.Email = email
	c.Phone = phone
	return c
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
