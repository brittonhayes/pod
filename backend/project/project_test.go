package project_test

import (
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/backend/config"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	config.Run()

	name := "Example Project"
	summary := "Example summary"
	client := "Example client"

	p := project.NewProject(filepath.Join(t.TempDir(), "pod.db"))
	p.With(name, summary, client)

	t.Run("create new project", func(t *testing.T) {
		assert.Equal(t, name, p.Name)
		assert.Equal(t, summary, p.Summary)
	})

	t.Run("save a project", func(t *testing.T) {
		err := p.Save()
		assert.NoError(t, err)
	})

	t.Run("query a project by name", func(t *testing.T) {
		projects, err := p.FindByName(name[0:3], 10)
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
		err := p.Delete(name)
		assert.NoError(t, err)
	})
}
