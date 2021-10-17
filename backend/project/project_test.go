package project_test

import (
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/config"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	config.Setup()

	defaultName := "Example Project"
	defaultSummary := "Example summary"
	defaultClient := &client.Client{
		ID:          1,
		Name:        "Joey",
		Description: "some words about joey",
		Email:       "fake@example.com",
		Phone:       "123-456-7890",
	}

	dbPath := filepath.Join(t.TempDir(), "pod.db")

	p := project.New(dbPath)
	p.With(defaultName, defaultSummary, defaultClient)

	c := client.New(dbPath)
	c.With(defaultClient.Name, defaultClient.Description, defaultClient.Email, defaultClient.Phone)
	c.Save()

	t.Run("create new project", func(t *testing.T) {
		assert.Equal(t, defaultName, p.Name)
		assert.Equal(t, defaultSummary, p.Summary)
	})

	t.Run("save a project", func(t *testing.T) {
		err := p.Save()
		assert.NoError(t, err)
	})

	t.Run("save a project from JSON", func(t *testing.T) {
		t.Logf("%#v", p)
		err := p.SaveJSON(map[string]interface{}{
			"name":    defaultName,
			"summary": defaultSummary,
			"Client":  defaultClient,
		})
		assert.NoError(t, err)
		defer p.Delete(defaultName)
	})

	t.Run("query a project by name", func(t *testing.T) {
		projects, err := p.FindByName(defaultName[0:3], 10)
		if assert.NoError(t, err) {
			assert.NotNil(t, projects)
		}
	})

	t.Run("list projects", func(t *testing.T) {
		projects, err := p.List()
		if assert.NoError(t, err) {
			assert.NotZero(t, len(projects))
		}
	})

	t.Run("delete a project", func(t *testing.T) {
		err := p.Delete(defaultName)
		assert.NoError(t, err)
	})

}
