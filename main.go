package main

import (
	_ "embed"

	"github.com/brittonhayes/pod/backend/bind"
	"github.com/brittonhayes/pod/backend/config"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	err := config.Setup()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1400,
		Height:    900,
		Title:     "pod",
		JS:        js,
		CSS:       css,
		Resizable: true,
		Colour:    "#131313",
	})

	cfg := config.NewConfig()

	app.Bind(cfg)
	app.Bind(&bind.Storage{})
	app.Bind(&bind.ProjectState{})

	err = app.Run()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
