package project

import (
	"encoding/json"
	"strings"

	"github.com/brittonhayes/pod/backend/store"
)

type Storage struct{}

func (s *Storage) SaveClient(payload string) (*Client, error) {
	client := NewClient()
	err := json.Unmarshal([]byte(payload), &client)
	if err != nil {
		return client, err
	}

	_, err = client.Save()
	if err != nil {
		return client, err
	}

	return client, nil
}

func (s *Storage) QueryClients(field, value string) ([]Client, error) {
	client := NewClient()
	results := []Client{}
	_, err := client.Query(strings.ToTitle(field), value, results)
	if err != nil {
		return []Client{}, err
	}

	return results, nil
}

func (s *Storage) ListProjects() ([]Project, error) {
	db, err := store.NewDB(DBName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results := []Project{}
	err = db.List(&results)
	return results, err
}

func (s *Storage) QueryProjects(field, value string) ([]Project, error) {
	project := NewProject()
	results := []Project{}
	_, err := project.Query(strings.ToTitle(field), value, results)
	if err != nil {
		return []Project{}, err
	}

	return results, nil
}

func (s *Storage) SaveProject(payload string) (*Project, error) {
	project := NewProject()
	err := json.Unmarshal([]byte(payload), &project)
	if err != nil {
		return project, err
	}

	_, err = project.Save()
	if err != nil {
		return project, err
	}

	return project, nil
}
