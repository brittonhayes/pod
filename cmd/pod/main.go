package main

import (
	_ "embed"
	"os"

	"github.com/brittonhayes/pod/internal/tray"
	"github.com/getlantern/systray"
	"github.com/rs/zerolog/log"
)

//go:embed logo.png
var logo []byte

func init() {
	err := tray.Initialize()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func onExit() {
	log.Info().Msg("exiting application")
	os.Exit(0)
}

func onReady() {
	systray.SetIcon(logo)
	systray.SetTitle("pod")
	systray.SetTooltip("pro audio project utility")

	projects := systray.AddMenuItem("View Projects", "View projects in files")
	go tray.ViewProjects(projects)

	clients := systray.AddMenuItem("View Clients", "View clients in files")
	go tray.ViewClients(clients)

	systray.AddSeparator()

	quitBtn := systray.AddMenuItem("Quit", "Quit the application")
	go tray.Quit(quitBtn)
}

func main() {
	systray.Run(onReady, onExit)
}
