package config_test

import (
	"testing"

	"github.com/brittonhayes/pod/backend/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("test new config", func(t *testing.T) {
		c := config.NewConfig()
		assert.NotNil(t, c)
	})

	t.Run("test config directory", func(t *testing.T) {
		dir := config.Directory()
		assert.Contains(t, dir, "pod")
	})

	t.Run("test config setup", func(t *testing.T) {
		c := config.Setup()
		assert.NoError(t, c)
		assert.DirExists(t, config.Directory())
	})
}
