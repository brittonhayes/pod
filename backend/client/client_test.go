package client_test

import (
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/config"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	config.Setup()

	defaultName := "Example Client"
	defaultDescription := "Example description"
	defaultEmail := "foo@email.lan"
	defaultPhone := "123-456-7890"

	c := client.New(filepath.Join(t.TempDir(), "pod.db"))
	c.With(defaultName, defaultDescription, defaultEmail, defaultPhone)

	t.Run("create new client", func(t *testing.T) {
		assert.Equal(t, defaultName, c.Name)
		assert.Equal(t, defaultDescription, c.Description)
	})

	t.Run("save a client", func(t *testing.T) {
		err := c.Save()
		assert.NoError(t, err)
	})

	t.Run("query a client by name", func(t *testing.T) {
		clients, err := c.FindByName(defaultName[0:3], 10)
		if assert.NoError(t, err) {
			assert.NotNil(t, clients)
		}
	})

	t.Run("list clients", func(t *testing.T) {
		clients, err := c.List()
		if assert.NoError(t, err) {
			assert.NotZero(t, len(clients))
		}
	})

	t.Run("delete a client", func(t *testing.T) {
		err := c.Delete(defaultName)
		assert.NoError(t, err)
	})
}
