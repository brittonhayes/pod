package project_test

import (
	"os"
	"testing"

	"github.com/brittonhayes/pod/backend/internal/system"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/brittonhayes/pod/backend/store"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	system.Initialize()

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

	defer t.Cleanup(func() {
		p := store.Path(dbName)
		t.Log(p)
		err := os.Remove(p)
		if err != nil {
			t.Fatal(err)
		}
	})
}
