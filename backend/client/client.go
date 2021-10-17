package client

import (
	"time"

	"github.com/brittonhayes/pod/backend/store"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

var (
	TableName = "clients"
)

type Social struct {
	Website   string `json:"website" mapstructure:"social_website"`
	Instagram string `json:"instagram" mapstructure:"social_instagram"`
	Twitter   string `json:"twitter"  mapstructure:"social_twitter"`
	Facebook  string `json:"facebook"  mapstructure:"social_facebook"`
	Linkedin  string `json:"linkedin"  mapstructure:"social_linkedin"`
}
type Client struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Social      Social `gorm:"embedded;embeddedPrefix:social_" mapstructure:",squash"`

	db string `gorm:"-"`
}

func New(path string) *Client {
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

func (c *Client) GetByID(id int) (*Client, error) {
	db, err := store.New(c.db, &Client{})
	if err != nil {
		return nil, err
	}

	tx := db.Find(c, id)
	return c, tx.Error
}

func (c *Client) Save() error {
	db, err := store.New(c.db, &Client{})
	if err != nil {
		return err
	}

	tx := db.Model(&Client{}).Create(c)

	return tx.Error
}

func (c *Client) Update(payload map[string]interface{}) error {
	db, err := store.New(c.db, &Client{})
	if err != nil {
		return err
	}

	err = mapstructure.Decode(payload, &c)
	if err != nil {
		return err
	}

	tx := db.Model(&c).Updates(c)

	return tx.Error
}

func (c *Client) SaveJSON(payload map[string]interface{}) error {
	db, err := store.New(c.db, &Client{})
	if err != nil {
		return err
	}

	err = mapstructure.Decode(payload, &c)
	if err != nil {
		return err
	}

	tx := db.Model(&Client{}).Create(c)
	return tx.Error
}

func (c *Client) Delete(name string) error {
	db, err := store.New(c.db)
	if err != nil {
		return err
	}
	tx := db.Where("name = ?", name).Delete(c)
	return tx.Error
}

func (c *Client) FindByName(name string, limit int) ([]Client, error) {
	db, err := store.New(c.db, &Client{})
	if err != nil {
		return nil, err
	}

	var clients []Client
	tx := db.Model(c).Limit(limit).Where("name LIKE ?", name).Find(&clients)

	return clients, tx.Error
}

func (c *Client) List() ([]Client, error) {
	db, err := store.New(c.db, &Client{})
	if err != nil {
		return nil, err
	}

	var projects []Client
	tx := db.Table(TableName).Select("*").Scan(&projects)
	return projects, tx.Error
}
