package project_test

import (
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

	p := project.NewProject()
	p.With(name, summary, client)

	t.Run("create new project", func(t *testing.T) {
		assert.Equal(t, name, p.Name)
		assert.Equal(t, summary, p.Summary)
	})

	t.Run("save a project", func(t *testing.T) {
		ok, err := p.Save()
		if assert.NoError(t, err) {
			assert.True(t, ok)
		}
	})

	t.Run("query a project", func(t *testing.T) {
		var projects []project.Project

		ok, err := p.Query("Name", name[0:3], projects)
		if assert.NoError(t, err) {
			assert.True(t, ok)
		}
	})

	t.Run("list projects", func(t *testing.T) {
		var projects []project.Project

		ok, err := p.List(projects)
		if assert.NoError(t, err) {
			assert.True(t, ok)
			assert.Greater(t, 1, len(projects))
		}
	})

	t.Run("delete a project", func(t *testing.T) {
		ok, err := p.Delete()
		if assert.NoError(t, err) {
			assert.True(t, ok)
		}
	})
}
