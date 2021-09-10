package project_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/project"
	"github.com/brittonhayes/pod/store"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	dbName := "pod"

	arg := "Example Client"
	c := project.NewClient(arg)

	t.Run("create new client", func(t *testing.T) {
		assert.Equal(t, arg, c.Name)
	})

	t.Run("save a client", func(t *testing.T) {
		ok, err := c.Save()
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("query a client", func(t *testing.T) {
		var clients []project.Client

		ok, err := c.Query("Name", arg[0:3], clients)
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("delete a client", func(t *testing.T) {
		ok, err := c.Delete()
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Cleanup(func() {
		path := filepath.Dir(store.Path(dbName))
		os.Remove(path)
	})
}
