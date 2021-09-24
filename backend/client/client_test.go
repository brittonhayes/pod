package client_test

import (
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	config.Run()

	c := client.NewClient(filepath.Join(t.TempDir(), "pod.db"))

	name := "Example Client"
	description := "Example description"
	c.With(name, description, "email", "phone")
	t.Logf("Testing with client: %#v\n", c)

	t.Run("create new client", func(t *testing.T) {
		assert.Equal(t, name, c.Name)
	})

	t.Run("save a client", func(t *testing.T) {
		ok, err := c.Save()
		if assert.NoError(t, err) {
			assert.True(t, ok)
		}
	})

	t.Run("query a client", func(t *testing.T) {
		var clients []client.Client

		ok, err := c.Query("Name", name[0:3], clients)
		if assert.NoError(t, err) {
			assert.True(t, ok)
		}

		t.Log(clients)
	})

	t.Run("list clients", func(t *testing.T) {
		var clients []client.Client

		ok, err := c.Query("Name", name[0:3], clients)
		if assert.NoError(t, err) {
			assert.True(t, ok)
		}
	})

	t.Run("delete a client", func(t *testing.T) {
		ok, err := c.Delete()
		if assert.NoError(t, err) {
			assert.True(t, ok)
		}
	})
}
