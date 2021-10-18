package store_test

import (
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/brittonhayes/pod/backend/store"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	t.Run("parse keywords", func(t *testing.T) {
		key := "key"
		args := map[string]string{
			"foo":  "key:foo",
			"baz":  "key:baz",
			"bing": "key:bing",
		}

		for expect, raw := range args {
			assert.Equal(t, expect, store.ParseKeywords(key, raw))
		}
	})

	t.Run("test new db", func(t *testing.T) {
		dbPath := filepath.Join(t.TempDir(), "test.db")

		db, err := store.New(dbPath, &project.Project{
			Client: &client.Client{
				Name: "johnny",
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, db)
	})
}
