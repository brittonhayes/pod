package store_test

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/brittonhayes/pod/backend/store"
	"github.com/stretchr/testify/assert"
)

type Example struct {
	ID   int    `storm:"id,increment"`
	Name string `storm:"index"`
}

func TestDB(t *testing.T) {

	db, err := store.NewDB(filepath.Join(t.TempDir(), "projects.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
		if assert.NoError(t, err) {
			assert.Equal(t, &arg, &to)
		}

		err = db.Get("ID", 1, &to)
		if assert.NoError(t, err) {
			assert.Equal(t, &arg, &to)
		}
	})

	t.Run("test store list", func(t *testing.T) {
		to := []Example{}
		want := defaultName
		err := db.List(&to)

		if assert.NoError(t, err) {
			assert.Equal(t, want, to[0].Name)
		}

		t.Logf("store list: %#v", &to)
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
		if assert.NoError(t, err) {
			assert.Equal(t, expect, to)
		}
	})

	t.Run("test store delete", func(t *testing.T) {
		err := db.Delete(&arg)
		assert.NoError(t, err)
	})
}
