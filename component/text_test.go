package component_test

import (
	"testing"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"github.com/brittonhayes/pod/component"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	canvas := test.NewCanvas()

	t.Run("test new text", func(t *testing.T) {
		expect := component.NewText("Hello")

		canvas.SetContent(expect)
		assert.Equal(t, expect, canvas.Content())
	})

	t.Run("test new text bind", func(t *testing.T) {
		b := binding.NewString()
		_, err := component.NewTextBind("Hello", b)
		assert.NoError(t, err)
	})
}
