package palette_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/brittonhayes/pod/palette"
)

func TestTheme(t *testing.T) {
	nord := &palette.Nord{}
	test.ApplyTheme(t, nord)
}
