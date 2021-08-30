package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/brittonhayes/pod/palette"
)

func NewText(text string) fyne.CanvasObject {
	nord := &palette.Nord{}
	return canvas.NewText(text, nord.Color(theme.ColorNameForeground, theme.VariantDark))
}

func NewTextBind(text string, bind binding.String) (*widget.Label, error) {
	err := bind.Set(text)
	if err != nil {
		return nil, err
	}

	return widget.NewLabelWithData(bind), nil
}
