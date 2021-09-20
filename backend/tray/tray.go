package tray

import (
	"fmt"

	"github.com/brittonhayes/pod/backend/client"
	"github.com/brittonhayes/pod/backend/project"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/rs/zerolog/log"
	"github.com/skratchdot/open-golang/open"
)

var (
	ErrTitle = "Pod"
)

type Handler interface {
	Handle(m *systray.MenuItem)
}

func errNotify(title, body string) {
	err := beeep.Notify(title, body, "")
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
}

func NewMenuButton(title, tooltip string, handler func(m *systray.MenuItem)) {
	item := systray.AddMenuItem(title, tooltip)
	go func() {
		<-item.ClickedCh
		handler(item)
	}()
}

func ViewProjects(m *systray.MenuItem) {
	<-m.ClickedCh
	log.Info().Msg("clicked view projects")
	path := project.Folder()
	err := open.Run(path)
	if err != nil {
		log.Error().Err(err).Send()
		errNotify(ErrTitle, fmt.Sprintf("couldn't find %s", path))
		return
	}
}

func ViewClients(m *systray.MenuItem) {
	<-m.ClickedCh
	log.Info().Msg("clicked view clients")
	path := client.Folder()
	err := open.Run(path)
	if err != nil {
		log.Error().Err(err).Send()
		errNotify(ErrTitle, fmt.Sprintf("couldn't find %s", path))
		return
	}
}

func Quit(m *systray.MenuItem) {
	<-m.ClickedCh
	log.Info().Msg("requested application quit")
	systray.Quit()
}
