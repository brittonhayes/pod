package store_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/store"
	"github.com/stretchr/testify/assert"
)

type Example struct {
	ID   int    `storm:"id,increment"`
	Name string `storm:"index"`
}

func TestDB(t *testing.T) {
	dbName := "pod"

	db, err := store.NewDB(dbName)
	if err != nil {
		log.Fatal(err)
	}

	defaultField := "Name"
	defaultName := "test_name"
	arg := Example{
		Name: defaultName,
	}

	t.Run("test store set", func(t *testing.T) {
		err := db.Set(&arg)
		assert.NoError(t, err)
	})

	t.Run("test store get", func(t *testing.T) {
		to := Example{}

		err := db.Get(defaultField, defaultName, &to)
		assert.NoError(t, err)

		assert.Equal(t, &arg, &to)
	})

	t.Run("test store query", func(t *testing.T) {
		var to []Example
		expect := []Example{
			{
				ID:   1,
				Name: defaultName,
			},
		}

		err := db.Query(defaultField, defaultName[0:3], &to)
		assert.NoError(t, err)

		assert.Equal(t, expect, to)
	})

	t.Run("test store delete", func(t *testing.T) {
		err := db.Delete(&arg)
		assert.NoError(t, err)
	})

	t.Cleanup(func() {
		path := filepath.Dir(store.Path(dbName))
		os.Remove(path)
	})
}
