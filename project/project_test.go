package project_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/project"
	"github.com/brittonhayes/pod/store"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	dbName := "pod"

	arg := "Example Project"
	p := project.NewProject(arg)

	t.Run("create new project", func(t *testing.T) {
		assert.Equal(t, arg, p.Name)
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

	t.Run("delete a project", func(t *testing.T) {
		ok, err := p.Delete()
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Cleanup(func() {
		path := filepath.Dir(store.Path(dbName))
		os.Remove(path)
	})

}
