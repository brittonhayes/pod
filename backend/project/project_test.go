package project_test

import (
	"os"
	"testing"

	"github.com/brittonhayes/pod/backend/config"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/brittonhayes/pod/backend/store"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	config.Run()

	dbName := "pod"

	arg := "Example Project"
	p := project.NewProject()
	p.With(arg, arg, arg)

	t.Run("create new project", func(t *testing.T) {
		assert.Equal(t, arg, p.Name)
		assert.Equal(t, arg, p.Summary)
	})

	t.Run("save a project", func(t *testing.T) {
		ok, err := p.Save()
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("query a project", func(t *testing.T) {
		var projects []project.Project

		ok, err := p.Query("Name", arg[0:3], projects)
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("list projects", func(t *testing.T) {
		var projects []*project.Project

		ok, err := p.List(projects)
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("delete a project", func(t *testing.T) {
		ok, err := p.Delete()
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	defer t.Cleanup(func() {
		p := store.Path(dbName)
		t.Log(p)
		err := os.Remove(p)
		if err != nil {
			t.Fatal(err)
		}
	})
}
