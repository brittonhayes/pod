package main

import (
	_ "embed"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/rs/zerolog/log"
)

const (
	ExitStatusOK = 0
)

//go:embed logo.png
var logo []byte

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
	log.Info().Msg("Exiting application")
	os.Exit(ExitStatusOK)
}

func handleQuit(m *systray.MenuItem) {
	<-m.ClickedCh
	log.Info().Msg("Requested application quit")
	systray.Quit()
}

func onReady() {
	systray.SetIcon(logo)
	systray.SetTitle("pod")
	systray.SetTooltip("pro audio project utility")
	mQuit := systray.AddMenuItem("Quit", "Quit the pod application")
	go handleQuit(mQuit)

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)
}
