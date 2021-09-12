package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	app := wails.CreateApp(&wails.AppConfig{
		Width:     1400,
		Height:    900,
		Title:     "pod",
		JS:        js,
		CSS:       css,
		Resizable: true,
		Colour:    "#131313",
	})
	app.Run()
}
