package store_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/brittonhayes/pod/store"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {

	key := "example"
	arg := "hello"
	expect := "hello"

	app := test.NewApp()
	s := store.New(app)

	t.Run("test set value", func(t *testing.T) {
		s.Set(key, arg)

		got := app.Preferences().String(key)
		assert.Equal(t, expect, got)
	})

	t.Run("test get value", func(t *testing.T) {

		got := s.Get(key)
		assert.Equal(t, expect, got)
	})

}
