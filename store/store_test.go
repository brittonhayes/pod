package store_test

import (
	"log"
	"os"
	"testing"

	"github.com/brittonhayes/pod/store"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	dbName := "pod"

	db, err := store.NewDB(dbName)
	if err != nil {
		log.Fatal(err)
	}

	defaultKey := "test_key"
	defaultValue := map[string]interface{}{
		"test": "test",
	}

	t.Run("test store set", func(t *testing.T) {
		err := db.Set(defaultKey, &defaultValue)
		assert.NoError(t, err)
	})

	t.Run("test store get", func(t *testing.T) {
		value := map[string]interface{}{}

		err := db.Get(defaultKey, &value)
		assert.NoError(t, err)

		assert.Equal(t, value, defaultValue)
	})

	t.Run("test store delete", func(t *testing.T) {
		value := map[string]interface{}{}

		err := db.Delete(defaultKey)
		assert.NoError(t, err)

		assert.Empty(t, value)
	})

	t.Cleanup(func() {
		path, err := db.Path()
		if err != nil {
			log.Fatal(err)
		}

		os.Remove(path)
	})
}
