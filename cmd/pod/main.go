package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/brittonhayes/pod/component"
	"github.com/brittonhayes/pod/palette"
	"github.com/brittonhayes/pod/store"
)

func main() {
	a := app.NewWithID("xyz.brittonhayes.pod")
	nord := palette.Nord{}

	db := store.New(a)

	w := a.NewWindow("Container")

	text1 := component.NewText("Hello")
	text2 := component.NewText("Hello")

	button := widget.NewButtonWithIcon("Click me 1", a.Icon(), func() { log.Println("clicked") })
	button2 := widget.NewButtonWithIcon("Click me 2", a.Icon(), func() {
		db.Set("greeting", "hello")
		greeting := db.Get("greeting")
		a.SendNotification(fyne.NewNotification("Updated in place", greeting))
	})

	form := component.NewForm(
		component.NewField(text1, button),
		component.NewField(text2, button2),
	)

	a.Settings().SetTheme(&nord)
	w.SetContent(form)
	w.ShowAndRun()
}
