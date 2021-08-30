package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type FormField struct {
	Label fyne.CanvasObject
	Input fyne.CanvasObject
}

func NewField(label, input fyne.CanvasObject) FormField {
	return FormField{
		Label: label,
		Input: input,
	}
}

func NewForm(fields ...FormField) *fyne.Container {
	ff := []fyne.CanvasObject{}
	for _, field := range fields {
		ff = append(ff, field.Label, field.Input)
	}

	return container.New(layout.NewFormLayout(), ff...)
}
